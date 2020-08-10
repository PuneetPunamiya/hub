import axios from "axios";
import { API_URL } from "../config/constants";
import { ICategory } from "../store/category";
import { IResource } from "../store/resources";
export interface Api {
  categories(): Promise<ICategory>;
  resources(): Promise<IResource>;
}
export class Hub implements Api {
  async resources() {
    try {
      return axios.get(
        `https://api-tekton-hub-staging.apps.openshift-web.p0s5.p1.openshiftapps.com/resources`
      );
    } catch (err) {
      return err.response;
    }
  }
  async categories() {
    try {
      return axios.get(`${API_URL}/categories`);
    } catch (err) {
      return err.response;
    }
  }
}
