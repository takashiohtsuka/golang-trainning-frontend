export interface WomanStore {
  id: number;
  name: string;
  business_type: string;
}

export interface WomanImage {
  id: number;
  path: string;
}

export interface WomanBlog {
  id: number;
  title: string;
}

export interface WomanListItem {
  id: number;
  name: string;
  age: number | null;
  birthplace: string | null;
  blood_type: string | null;
  hobby: string | null;
  stores: WomanStore[];
  images: WomanImage[];
  blogs: WomanBlog[];
}

export interface DistrictWomenResponse {
  women: WomanListItem[];
  total: number;
}
