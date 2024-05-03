import type { AxiosError, AxiosResponse } from "axios";
import axios from "axios";

type HTTPMethod = "GET" | "POST" | "PUT" | "DELETE";

export type APIResponse<T> = {
  ok: boolean;
  status: number;
  data?: T;
  error?: AxiosError;
};

const getHeaders = ({ accessToken }: { accessToken?: string }) => {
  return {
    // Authorization: `Bearer ${accessToken}`,
    // add your headers here
    "Access-Control-Allow-Origin": "*",
    // "Access-Control-Allow-Methods": "GET, POST, PUT, DELETE",
  };
};

export const httpRequest = async <T, D>(
  url: string,
  method: HTTPMethod = "GET",
  data?: D,
  options?: {
    headers?: Record<string, string>;
    accessToken?: string;
  },
): Promise<APIResponse<T>> => {
  try {
    const customHeaders = getHeaders({
      accessToken: options?.accessToken,
    });
    const response: AxiosResponse<T> = await axios({
      method,
      url,
      headers: {
        ...options?.headers,
        ...customHeaders,
      },
      data,
    });

    return {
      ok: true,
      status: response.status,
      data: response.data,
    };
  } catch (error) {
    return {
      ok: false,
      status: (error as AxiosError).response?.status ?? 500,
      error: error as AxiosError,
    };
  }
};