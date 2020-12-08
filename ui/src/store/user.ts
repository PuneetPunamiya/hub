import { flow, getEnv, Instance, types } from 'mobx-state-tree';
import { Api } from '../api';

export const TokenInfo = types.model({
  token: types.optional(types.string, ''),
  expiresAt: types.optional(types.number, 0),
  refreshInterval: types.optional(types.string, '')
});

export type IAuthStore = Instance<typeof AuthStore>;
export type ITokenInfo = Instance<typeof TokenInfo>;
export interface AuthCodeProps {
  code: string;
}

export const AuthStore = types
  .model({
    accessTokenInfo: types.optional(TokenInfo, {}),
    refreshTokenInfo: types.optional(TokenInfo, {}),
    isLoading: false,
    isAuthenticated: false,
    err: ''
  })
  .actions((self) => ({
    addAccessTokenInfo(item: ITokenInfo) {
      self.accessTokenInfo.token = item.token;
      self.accessTokenInfo.expiresAt = item.expiresAt;
      self.accessTokenInfo.refreshInterval = item.refreshInterval;
    },
    addRefreshTokenInfo(item: ITokenInfo) {
      self.refreshTokenInfo.token = item.token;
      self.refreshTokenInfo.expiresAt = item.expiresAt;
      self.refreshTokenInfo.refreshInterval = item.refreshInterval;
    },
    setIsAuthenticated(l: boolean) {
      self.isAuthenticated = l;
    },
    setIsLoading(l: boolean) {
      self.isLoading = l;
    },
    onFailure(err: string) {
      self.err = err;
    }
  }))
  .views((self) => ({
    get api(): Api {
      return getEnv(self).api;
    }
  }))
  .actions((self) => ({
    authenticate: flow(function* (authCode: AuthCodeProps) {
      try {
        const { api } = self;

        const json = yield api.authentication(authCode.code);

        self.addAccessTokenInfo(json.data.access);
        self.addRefreshTokenInfo(json.data.refresh);

        self.setIsAuthenticated(true);
        self.setIsLoading(true);
      } catch (err) {
        self.err = err.toString();
      }
    })
  }));
