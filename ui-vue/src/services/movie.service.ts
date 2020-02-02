import api from "./api.service";
export const searchMovie = async (searchTerm: string) => {
  try {
    const res = await api.get(`/movie/search/${searchTerm}`);

    return res;
  } catch (error) {
    throw new Error(error.response.data);
  }
};
