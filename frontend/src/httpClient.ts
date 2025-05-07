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

  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`);
  }

  const json = await response.json();
  return {
    data: json.data || json,
    total: json.total,
  };
}; 