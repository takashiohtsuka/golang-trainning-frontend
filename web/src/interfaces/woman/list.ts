export interface StoreAssignment {
  id: number;
  store_id: number;
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
  store_assignments: StoreAssignment[];
  images: WomanImage[];
  blogs: WomanBlog[];
}

export interface DistrictWomenResponse {
  women: WomanListItem[];
}
