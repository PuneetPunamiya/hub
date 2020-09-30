import { RequestFactoryType, ApiResponse, MyApi } from './client';

const Hub: RequestFactoryType = (): Promise<ApiResponse<any>> => {
  console.log('This works');
  return new Promise<ApiResponse<any>>((resolve) => {
    resolve();
  });
};

// eslint-disable-next-line @typescript-eslint/ban-types
const api = new MyApi<{}>({}, Hub);

const data = api.GetCategories();
data.then((res) => console.log(res.json()));
