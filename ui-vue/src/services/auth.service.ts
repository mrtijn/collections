import api from "./api.service";
export const Login = async (email: string, password: string) => {
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

export const Register = async (form: any) => {
  try {
    const { token, expiresAt } = await api.post("/register", form);

    setToken(token, expiresAt);

    return true;
  } catch (error) {
    throw new Error(error.response.data);
  }
};

export const getUser = async () => {
  try {
    const res = await api.get("/user/55");

    console.log(res);
  } catch (error) {
    throw new Error(error.response.data);
  }
};

export const Logout = () => {
  window.localStorage.removeItem("token");
  window.localStorage.removeItem("expiresAt");
};

export const isLoggedIn = () => {
  let token = window.localStorage.getItem("token");

  if (token) {
    return true;
  }

  return false;
};

function setToken(token: string, expiresAt: string) {
  window.localStorage.setItem("token", token);
  window.localStorage.setItem("expiresAt", expiresAt);
}
