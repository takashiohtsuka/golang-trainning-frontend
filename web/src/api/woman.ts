import type { DistrictWomenResponse } from "@/interfaces/district/womanList";

const BASE_URL = "http://localhost:8081/frontend";

export async function fetchDistrictWomen(districtId: string): Promise<DistrictWomenResponse> {
  const res = await fetch(`${BASE_URL}/districts/${districtId}/women`);
  if (!res.ok) throw new Error(`${res.status}`);
  return res.json() as Promise<DistrictWomenResponse>;
}
