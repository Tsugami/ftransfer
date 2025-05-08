import type { DataProvider } from 'react-admin';
import { httpClient } from './httpClient';

const apiUrl = import.meta.env.VITE_API_URL ?? 'http://localhost:8080/api/v1';

export const dataProvider: DataProvider = {
  getList: async (resource, _params) => {
    const { data, total } = await httpClient(`${apiUrl}/${resource}`);
    return {
      data,
      total,
    };
  },

  getOne: async (resource, params) => {
    const { data } = await httpClient(`${apiUrl}/${resource}/${params.id}`);
    return {
      data,
    };
  },

  getMany: async (resource, params) => {
    const { data } = await httpClient(`${apiUrl}/${resource}?ids=${params.ids.join(',')}`);
    return {
      data,
    };
  },

  getManyReference: async (resource, params) => {
    const { data, total } = await httpClient(
      `${apiUrl}/${resource}?${params.target}=${params.id}`
    );
    return {
      data,
      total,
    };
  },

  create: async (resource, params) => {
    const { data } = await httpClient(`${apiUrl}/${resource}`, {
      method: 'POST',
      body: JSON.stringify(params.data),
    });
    return {
      data,
    };
  },

  update: async (resource, params) => {
    console.log('updating', params);
    const { data } = await httpClient(`${apiUrl}/${resource}/${params.id}`, {
      method: 'PUT',
      body: JSON.stringify(params.data),
    });
    return {
      data,
    };
  },

  updateMany: async (_resource, _params) => {
    throw new Error('Update many not supported');
  },

  delete: async (resource, params) => {
    const { data } = await httpClient(`${apiUrl}/${resource}/${params.id}`, {
      method: 'DELETE',
    });
    return {
      data,
    };
  },

  deleteMany: async (_resource, _params) => {
    throw new Error('Delete many not supported');
  },
}; 