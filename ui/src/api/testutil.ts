import { Api } from './';
import * as fs from 'fs';
import { ICategory } from '../store/category';
import { IResource, IVersion } from '../store/resource';
import { IAuthStore } from '../store/user';

export class FakeHub implements Api {
  dataDir: string;

  constructor(dataDir: string) {
    this.dataDir = dataDir;
  }

  async categories() {
    const data = `${this.dataDir}/categories.json`;
    const ret = () => JSON.parse(fs.readFileSync(data).toString());
    return new Promise<ICategory>((resolve) => {
      setTimeout(() => resolve(ret()), 1000);
    });
  }

  async resources() {
    const data = `${this.dataDir}/resources.json`;

    const ret = () => JSON.parse(fs.readFileSync(data).toString());
    return new Promise<IResource>((resolve) => {
      setTimeout(() => resolve(ret()), 1000);
    });
  }

  async resourceVersion(resourceId: string) {
    const data = `${this.dataDir}/resource_versions_${resourceId}.json`;

    const ret = () => JSON.parse(fs.readFileSync(data).toString());
    return new Promise<IVersion>((resolve) => {
      setTimeout(() => resolve(ret()), 1000);
    });
  }

  async versionUpdate(versionId: string) {
    const data = `${this.dataDir}/version_details_${versionId}.json`;

    const ret = () => JSON.parse(fs.readFileSync(data).toString());
    return new Promise<IVersion>((resolve) => {
      setTimeout(() => resolve(ret()), 1000);
    });
  }

  async authentication(authCode: string) {
    const data = `${this.dataDir}/authentication.json`;

    const ret = () => JSON.parse(fs.readFileSync(data).toString());
    return new Promise<IAuthStore>((resolve) => {
      setTimeout(() => resolve(ret()), 1000);
    });
  }
}
