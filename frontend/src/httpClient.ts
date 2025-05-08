interface RequestOptions {
  method?: string;
  body?: string;
  headers?: Record<string, string>;
}

interface Response {
  data: any;
  total?: number;
}

export const httpClient = async (url: string, options: RequestOptions = {}): Promise<Response> => {
  const defaultHeaders = {
    'Content-Type': 'application/json',
  };

  const response = await fetch(url, {
    ...options,
    headers: {
      ...defaultHeaders,
      ...options.headers,
    },
  });


  const json = response.status === 204 ? {} : await response?.json() ?? {};
  if (!response.ok) {
    const error = json?.error ?? response?.statusText;

    throw new Error(`Error! status: ${error}`);
  }


  return {
    data: json?.data || json,
    total: json?.total,
  };
}; 