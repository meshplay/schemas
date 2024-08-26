import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';

export const api = createApi({
  reducerPath: 'meshplayApi',
  baseQuery: fetchBaseQuery({ baseUrl: '/api/' }),
  endpoints: () => ({}),
});
