import api from "./api.service";
export const GetMyCollections = async () => {
  try {
    const res = await api.get(`/collections/self`);

    return res;
  } catch (error) {
    throw new Error(error.response.data);
  }
};

export const GetCollectionById = async (id: string) => {
  try {
    const res = await api.get(`/collection/${id}`);

    return res;
  } catch (error) {
    throw new Error(error.response.data);
  }
};
