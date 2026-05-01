export interface DistrictItem {
  id: number;
  name: string;
  prefecture_id: number;
}

export interface DistrictListResponse {
  districts: DistrictItem[];
}
