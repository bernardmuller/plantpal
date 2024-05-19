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
    "Access-Control-Allow-Origin": "*",
    "Access-Control-Allow-Credentials": "true",
    "Content-Type": "application/json",
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
    // const response: AxiosResponse<T> = await axios({
    //   method,
    //   url,
    //   headers: {
    //     ...options?.headers,
    //     ...customHeaders,
    //   },
    //   data,
    //   withCredentials: true,
    // });

    const response: Response = await fetch(url, {
      method,
      headers: {
        ...options?.headers,
        ...customHeaders,
      },
      body: data ? JSON.stringify(data) : undefined,
      credentials: 'same-origin',
    });


    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    const responseData: T = await response.json();

    return {
      ok: true,
      status: response.status,
      data: responseData,
    };
  } catch (error) {
    return {
      ok: false,
      status: (error as AxiosError).response?.status ?? 500,
      error: error as AxiosError,
    };
  }
};