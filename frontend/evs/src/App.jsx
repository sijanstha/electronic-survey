import AuthProvider from "./provider/authProvider";
import Routes from "./routes";
import { positions, Provider } from "react-alert";
import AlertTemplate from "react-alert-template-basic";

const options = {
  timeout: 3000,
  position: positions.BOTTOM_CENTER
};

const App = () => {
  return (
    <Provider template={AlertTemplate} {...options}>
      <AuthProvider>
        <Routes />
      </AuthProvider>
    </Provider>
  );
}

export default App;
