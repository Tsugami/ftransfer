import type { DataProvider } from 'react-admin';
import { httpClient } from './httpClient';

const apiUrl = 'http://localhost:8080/api/v1';

export const dataProvider: DataProvider = {
  getList: async (resource, params) => {
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
    const { data } = await httpClient(`${apiUrl}/${resource}/${params.id}`, {
      method: 'PUT',
      body: JSON.stringify(params.data),
    });
    return {
      data,
    };
  },

  updateMany: async (resource, params) => {
    const { data } = await httpClient(`${apiUrl}/${resource}`, {
      method: 'PUT',
      body: JSON.stringify(params.ids),
    });
    return {
      data,
    };
  },

  delete: async (resource, params) => {
    const { data } = await httpClient(`${apiUrl}/${resource}/${params.id}`, {
      method: 'DELETE',
    });
    return {
      data,
    };
  },

  deleteMany: async (resource, params) => {
    const { data } = await httpClient(`${apiUrl}/${resource}`, {
      method: 'DELETE',
      body: JSON.stringify(params.ids),
    });
    return {
      data,
    };
  },
}; 