export interface ImmediateAvailableWomanStore {
  id: number;
  name: string;
  business_type: string;
}

export interface ImmediateAvailableWomanImage {
  id: number;
  path: string;
}

export interface ImmediateAvailableWomanItem {
  id: number;
  name: string;
  age: number | null;
  birthplace: string | null;
  blood_type: string | null;
  hobby: string | null;
  store: ImmediateAvailableWomanStore;
  images: ImmediateAvailableWomanImage[];
  expires_at: string | null;
}

export interface ImmediateAvailableWomanListResponse {
  women: ImmediateAvailableWomanItem[];
  total: number;
}
