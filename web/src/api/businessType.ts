import { BASE_URL, fetchJSON } from "./base";
import type { BusinessTypeListResponse } from "@/interfaces/businessType/list";

export async function fetchBusinessTypes(options?: RequestInit): Promise<BusinessTypeListResponse> {
  return fetchJSON<BusinessTypeListResponse>(`${BASE_URL}/business_types`, options);
}
