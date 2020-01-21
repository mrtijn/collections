import api from "./api.service.ts";
export const Login = async (email, password) => {
  try {
    const { token, expiresAt } = await api.post("/login", {
      email,
      password
    });

    setToken(token, expiresAt);

    return true;
  } catch (error) {
    throw new Error(error.response.data);
  }
};

function setToken(token, expiresAt) {
  window.localStorage.setItem("token", token);
  window.localStorage.setItem("expiresAt", expiresAt);
}

export const isLoggedIn = () => {
  let token = window.localStorage.getItem("token");

  if (token) {
    return true;
  }

  return false;
};
