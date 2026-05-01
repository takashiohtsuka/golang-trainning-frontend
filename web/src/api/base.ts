export const BASE_URL = "http://localhost/frontend";

export async function fetchJSON<T>(url: string, options?: RequestInit): Promise<T> {
  const res = await fetch(url, options);
  if (!res.ok) throw new Error(`${res.status}`);
  return res.json() as Promise<T>;
}
