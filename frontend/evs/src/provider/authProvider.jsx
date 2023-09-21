import { createContext, useContext, useEffect, useMemo, useState } from "react";
import jwt_decode from "jwt-decode";

const AuthContext = createContext();

const AuthProvider = ({ children }) => {
  // State to hold the authentication token
  const [token, setToken_] = useState(localStorage.getItem("token"));
  const [loggedInUserInfo, setLoggedInUserInfo_] = useState({email: '', role: '', id: ''})

  // Function to set the authentication token
  const setToken = (newToken) => {
    setToken_(newToken);
  };

  const getLoggedInUserInfo = (token) => loggedInUserInfo;

  useEffect(() => {
    if (token) {
      // axiosInstance.defaults.headers.common["Authorization"] = token;
      localStorage.setItem("token", token);
      var decoded = jwt_decode(token);

      setLoggedInUserInfo_({
        email: decoded.email,
        id: parseInt(decoded.id),
        role: decoded.role
      });

    } else {
      // delete axiosInstance.defaults.headers.common["Authorization"];
      localStorage.removeItem("token");
    }
  }, [token]);

  // Memoized value of the authentication context
  const contextValue = useMemo(
    () => ({
      token,
      setToken,
      getLoggedInUserInfo,
    }),
    [token,getLoggedInUserInfo]
  );

  console.log(getLoggedInUserInfo());

  // Provide the authentication context to the children components
  return (
    <AuthContext.Provider value={contextValue}>{children}</AuthContext.Provider>
  );
};

export const useAuth = () => {
  return useContext(AuthContext);
};

export default AuthProvider;