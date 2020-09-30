export type AuthAuthenticateInternalErrorResponseBody = {
  fault?: boolean;
  id?: string;
  message?: string;
  name: string;
  temporary?: boolean;
  timeout?: boolean;
};

export type AuthAuthenticateInvalidCodeResponseBody = {
  fault?: boolean;
  id?: string;
  message?: string;
  name: string;
  temporary?: boolean;
  timeout?: boolean;
};

export type AuthAuthenticateInvalidScopesResponseBody = {
  fault?: boolean;
  id?: string;
  message?: string;
  name: string;
  temporary?: boolean;
  timeout?: boolean;
};

export type AuthAuthenticateInvalidTokenResponseBody = {
  fault?: boolean;
  id?: string;
  message?: string;
  name: string;
  temporary?: boolean;
  timeout?: boolean;
};

export type AuthAuthenticateResponseBody = {
  token: string;
};

export type CatalogResponse = {
  id: number;
  name?: string;
  type?: string;
};

export type CatalogResponseBody = {
  id: number;
  name?: string;
  type?: string;
};

export type CategoryListInternalErrorResponseBody = {
  fault?: boolean;
  id?: string;
  message?: string;
  name: string;
  temporary?: boolean;
  timeout?: boolean;
};

export type CategoryResponse = {
  id: number;
  name?: string;
  tags?: Array<TagResponse>;
};

export type RatingGetInternalErrorResponseBody = {
  fault?: boolean;
  id?: string;
  message?: string;
  name: string;
  temporary?: boolean;
  timeout?: boolean;
};

export type RatingGetInvalidScopesResponseBody = {
  fault?: boolean;
  id?: string;
  message?: string;
  name: string;
  temporary?: boolean;
  timeout?: boolean;
};

export type RatingGetInvalidTokenResponseBody = {
  fault?: boolean;
  id?: string;
  message?: string;
  name: string;
  temporary?: boolean;
  timeout?: boolean;
};

export type RatingGetNotFoundResponseBody = {
  fault?: boolean;
  id?: string;
  message?: string;
  name: string;
  temporary?: boolean;
  timeout?: boolean;
};

export type RatingGetResponseBody = {
  rating: number;
};

export type RatingUpdateInternalErrorResponseBody = {
  fault?: boolean;
  id?: string;
  message?: string;
  name: string;
  temporary?: boolean;
  timeout?: boolean;
};

export type RatingUpdateInvalidScopesResponseBody = {
  fault?: boolean;
  id?: string;
  message?: string;
  name: string;
  temporary?: boolean;
  timeout?: boolean;
};

export type RatingUpdateInvalidTokenResponseBody = {
  fault?: boolean;
  id?: string;
  message?: string;
  name: string;
  temporary?: boolean;
  timeout?: boolean;
};

export type RatingUpdateNotFoundResponseBody = {
  fault?: boolean;
  id?: string;
  message?: string;
  name: string;
  temporary?: boolean;
  timeout?: boolean;
};

export type RatingUpdateRequestBody = {
  rating: number;
};

export type ResourceByCatalogKindNameInternalErrorResponseBody = {
  fault?: boolean;
  id?: string;
  message?: string;
  name: string;
  temporary?: boolean;
  timeout?: boolean;
};

export type ResourceByCatalogKindNameNotFoundResponseBody = {
  fault?: boolean;
  id?: string;
  message?: string;
  name: string;
  temporary?: boolean;
  timeout?: boolean;
};

export type ResourceByCatalogKindNameVersionInternalErrorResponseBody = {
  fault?: boolean;
  id?: string;
  message?: string;
  name: string;
  temporary?: boolean;
  timeout?: boolean;
};

export type ResourceByCatalogKindNameVersionNotFoundResponseBody = {
  fault?: boolean;
  id?: string;
  message?: string;
  name: string;
  temporary?: boolean;
  timeout?: boolean;
};

export type ResourceByCatalogKindNameVersionResponseBody = {
  description?: string;
  displayName?: string;
  id: number;
  minPipelinesVersion?: string;
  rawURL?: string;
  resource?: ResourceResponseBodyInfo;
  updatedAt?: string;
  version?: string;
  webURL?: string;
};

export type ResourceByIDInternalErrorResponseBody = {
  fault?: boolean;
  id?: string;
  message?: string;
  name: string;
  temporary?: boolean;
  timeout?: boolean;
};

export type ResourceByIDNotFoundResponseBody = {
  fault?: boolean;
  id?: string;
  message?: string;
  name: string;
  temporary?: boolean;
  timeout?: boolean;
};

export type ResourceByIDResponseBody = {
  catalog?: CatalogResponseBody;
  id: number;
  kind?: string;
  latestVersion?: VersionResponseBodyWithoutResource;
  name?: string;
  rating?: number;
  tags?: Array<TagResponseBody>;
  versions?: Array<VersionResponseBodyTiny>;
};

export type ResourceByVersionIDInternalErrorResponseBody = {
  fault?: boolean;
  id?: string;
  message?: string;
  name: string;
  temporary?: boolean;
  timeout?: boolean;
};

export type ResourceByVersionIDNotFoundResponseBody = {
  fault?: boolean;
  id?: string;
  message?: string;
  name: string;
  temporary?: boolean;
  timeout?: boolean;
};

export type ResourceByVersionIDResponseBody = {
  description?: string;
  displayName?: string;
  id: number;
  minPipelinesVersion?: string;
  rawURL?: string;
  resource?: ResourceResponseBodyInfo;
  updatedAt?: string;
  version?: string;
  webURL?: string;
};

export type ResourceListInternalErrorResponseBody = {
  fault?: boolean;
  id?: string;
  message?: string;
  name: string;
  temporary?: boolean;
  timeout?: boolean;
};

export type ResourceQueryInternalErrorResponseBody = {
  fault?: boolean;
  id?: string;
  message?: string;
  name: string;
  temporary?: boolean;
  timeout?: boolean;
};

export type ResourceQueryNotFoundResponseBody = {
  fault?: boolean;
  id?: string;
  message?: string;
  name: string;
  temporary?: boolean;
  timeout?: boolean;
};

export type ResourceResourceResponseWithoutVersionCollection = Array<
  ResourceResponseWithoutVersion
>;

export type ResourceResponseBodyInfo = {
  catalog?: CatalogResponseBody;
  id: number;
  kind?: string;
  name?: string;
  rating?: number;
  tags?: Array<TagResponseBody>;
};

export type ResourceResponseWithoutVersion = {
  catalog?: CatalogResponse;
  id: number;
  kind?: string;
  latestVersion?: VersionResponseWithoutResource;
  name?: string;
  rating?: number;
  tags?: Array<TagResponse>;
};

export type ResourceVersionsByIDInternalErrorResponseBody = {
  fault?: boolean;
  id?: string;
  message?: string;
  name: string;
  temporary?: boolean;
  timeout?: boolean;
};

export type ResourceVersionsByIDNotFoundResponseBody = {
  fault?: boolean;
  id?: string;
  message?: string;
  name: string;
  temporary?: boolean;
  timeout?: boolean;
};

export type ResourceVersionsByIDResponseBody = {
  latest: VersionResponseBodyMin;
  versions?: Array<VersionResponseBodyMin>;
};

export type StatusStatusResponseBody = {
  status: string;
};

export type TagResponse = {
  id: number;
  name?: string;
};

export type TagResponseBody = {
  id: number;
  name?: string;
};

export type VersionResponseBodyMin = {
  id: number;
  rawURL?: string;
  version?: string;
  webURL?: string;
};

export type VersionResponseBodyTiny = {
  id: number;
  version?: string;
};

export type VersionResponseBodyWithoutResource = {
  description?: string;
  displayName?: string;
  id: number;
  minPipelinesVersion?: string;
  rawURL?: string;
  updatedAt?: string;
  version?: string;
  webURL?: string;
};

export type VersionResponseWithoutResource = {
  description?: string;
  displayName?: string;
  id: number;
  minPipelinesVersion?: string;
  rawURL?: string;
  updatedAt?: string;
  version?: string;
  webURL?: string;
};

export type PostAuthLoginQueryParameters = {
  code: string;
};

export type GetQueryQueryParameters = {
  name?: string;
  kinds?: Array<string>;
  tags?: Array<string>;
  limit?: number;
  match?: string;
};

export type GetResourceByIdRatingHeaderParameters = {
  Authorization: any;
};

export type PutResourceByIdRatingBodyParameters = RatingUpdateRequestBody;

export type PutResourceByIdRatingHeaderParameters = {
  Authorization: any;
};

export type GetResourcesQueryParameters = {
  limit?: number;
};

export interface ApiResponse<T> extends Response {
  json(): Promise<T>;
}
export type RequestFactoryType = (
  path: string,
  query: any,
  body: any,
  formData: any,
  headers: any,
  method: string,
  configuration: any
) => Promise<ApiResponse<any>>;

export class MyApi<T extends {} = {}> {
  constructor(protected configuration: T, protected requestFactory: RequestFactoryType) {}
  Get(): Promise<ApiResponse<StatusStatusResponseBody>> {
    const path = '/';
    return this.requestFactory(
      path,
      undefined,
      undefined,
      undefined,
      undefined,
      'GET',
      this.configuration
    );
  }

  PostAuthLogin(
    query: PostAuthLoginQueryParameters
  ): Promise<
    ApiResponse<
      | AuthAuthenticateResponseBody
      | AuthAuthenticateInvalidCodeResponseBody
      | AuthAuthenticateInvalidTokenResponseBody
      | AuthAuthenticateInvalidScopesResponseBody
      | AuthAuthenticateInternalErrorResponseBody
    >
  > {
    const path = '/auth/login';
    return this.requestFactory(
      path,
      query,
      undefined,
      undefined,
      undefined,
      'POST',
      this.configuration
    );
  }

  GetCategories(): Promise<
    ApiResponse<Array<CategoryResponse> | CategoryListInternalErrorResponseBody>
  > {
    const path = '/categories';
    return this.requestFactory(
      path,
      undefined,
      undefined,
      undefined,
      undefined,
      'GET',
      this.configuration
    );
  }

  GetQuery(
    query: GetQueryQueryParameters
  ): Promise<
    ApiResponse<
      | ResourceResourceResponseWithoutVersionCollection
      | ResourceQueryNotFoundResponseBody
      | ResourceQueryInternalErrorResponseBody
    >
  > {
    const path = '/query';
    return this.requestFactory(
      path,
      query,
      undefined,
      undefined,
      undefined,
      'GET',
      this.configuration
    );
  }

  GetResourceVersionByVersionID(
    versionIDPathParameter: number
  ): Promise<
    ApiResponse<
      | ResourceByVersionIDResponseBody
      | ResourceByVersionIDNotFoundResponseBody
      | ResourceByVersionIDInternalErrorResponseBody
    >
  > {
    let path = '/resource/version/{versionID}';
    path = path.replace('{versionID}', String(versionIDPathParameter));
    return this.requestFactory(
      path,
      undefined,
      undefined,
      undefined,
      undefined,
      'GET',
      this.configuration
    );
  }

  GetResourceByCatalogByKindByName(
    catalogPathParameter: string,
    kindPathParameter: string,
    namePathParameter: string
  ): Promise<
    ApiResponse<
      | ResourceResourceResponseWithoutVersionCollection
      | ResourceByCatalogKindNameNotFoundResponseBody
      | ResourceByCatalogKindNameInternalErrorResponseBody
    >
  > {
    let path = '/resource/{catalog}/{kind}/{name}';
    path = path.replace('{catalog}', String(catalogPathParameter));

    path = path.replace('{kind}', String(kindPathParameter));

    path = path.replace('{name}', String(namePathParameter));
    return this.requestFactory(
      path,
      undefined,
      undefined,
      undefined,
      undefined,
      'GET',
      this.configuration
    );
  }

  GetResourceByCatalogByKindByNameByVersion(
    catalogPathParameter: string,
    kindPathParameter: string,
    namePathParameter: string,
    versionPathParameter: string
  ): Promise<
    ApiResponse<
      | ResourceByCatalogKindNameVersionResponseBody
      | ResourceByCatalogKindNameVersionNotFoundResponseBody
      | ResourceByCatalogKindNameVersionInternalErrorResponseBody
    >
  > {
    let path = '/resource/{catalog}/{kind}/{name}/{version}';
    path = path.replace('{catalog}', String(catalogPathParameter));

    path = path.replace('{kind}', String(kindPathParameter));

    path = path.replace('{name}', String(namePathParameter));

    path = path.replace('{version}', String(versionPathParameter));
    return this.requestFactory(
      path,
      undefined,
      undefined,
      undefined,
      undefined,
      'GET',
      this.configuration
    );
  }

  GetResourceById(
    idPathParameter: number
  ): Promise<
    ApiResponse<
      | ResourceByIDResponseBody
      | ResourceByIDNotFoundResponseBody
      | ResourceByIDInternalErrorResponseBody
    >
  > {
    let path = '/resource/{id}';
    path = path.replace('{id}', String(idPathParameter));
    return this.requestFactory(
      path,
      undefined,
      undefined,
      undefined,
      undefined,
      'GET',
      this.configuration
    );
  }

  GetResourceByIdRating(
    idPathParameter: number,
    header: GetResourceByIdRatingHeaderParameters
  ): Promise<
    ApiResponse<
      | RatingGetResponseBody
      | RatingGetInvalidTokenResponseBody
      | RatingGetInvalidScopesResponseBody
      | RatingGetNotFoundResponseBody
      | RatingGetInternalErrorResponseBody
    >
  > {
    let path = '/resource/{id}/rating';
    path = path.replace('{id}', String(idPathParameter));
    return this.requestFactory(
      path,
      undefined,
      undefined,
      undefined,
      header,
      'GET',
      this.configuration
    );
  }

  PutResourceByIdRating(
    idPathParameter: number,
    body: PutResourceByIdRatingBodyParameters,
    header: PutResourceByIdRatingHeaderParameters
  ): Promise<
    ApiResponse<
      | any
      | RatingUpdateInvalidTokenResponseBody
      | RatingUpdateInvalidScopesResponseBody
      | RatingUpdateNotFoundResponseBody
      | RatingUpdateInternalErrorResponseBody
    >
  > {
    let path = '/resource/{id}/rating';
    path = path.replace('{id}', String(idPathParameter));
    return this.requestFactory(path, undefined, body, undefined, header, 'PUT', this.configuration);
  }

  GetResourceByIdVersions(
    idPathParameter: number
  ): Promise<
    ApiResponse<
      | ResourceVersionsByIDResponseBody
      | ResourceVersionsByIDNotFoundResponseBody
      | ResourceVersionsByIDInternalErrorResponseBody
    >
  > {
    let path = '/resource/{id}/versions';
    path = path.replace('{id}', String(idPathParameter));
    return this.requestFactory(
      path,
      undefined,
      undefined,
      undefined,
      undefined,
      'GET',
      this.configuration
    );
  }

  GetResources(
    query: GetResourcesQueryParameters
  ): Promise<
    ApiResponse<
      ResourceResourceResponseWithoutVersionCollection | ResourceListInternalErrorResponseBody
    >
  > {
    const path = '/resources';
    return this.requestFactory(
      path,
      query,
      undefined,
      undefined,
      undefined,
      'GET',
      this.configuration
    );
  }

  GetSchemaSwaggerJson(): Promise<ApiResponse<any>> {
    const path = '/schema/swagger.json';
    return this.requestFactory(
      path,
      undefined,
      undefined,
      undefined,
      undefined,
      'GET',
      this.configuration
    );
  }
}
