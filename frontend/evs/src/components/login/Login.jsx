import { useState } from "react";
import "./login.css";
import {
  isContainWhiteSpace,
  isEmail,
  isEmpty,
  isLength,
} from "../../shared/validator";
import { axiosInstance } from "../../axiosConfig";
import { useAuth } from "../../provider/authProvider";
import { useNavigate } from "react-router-dom";

const Login = () => {
  const { setToken } = useAuth();
  const navigate = useNavigate();

  const [formState, setFormState] = useState({
    formData: { email: '', password: "" },
    errors: { email: "", password: "" }
  });

  const validateLoginForm = () => {
    let errors = {};
    const { formData } = formState;

    if (isEmpty(formData.email)) {
      errors.email = "Email can't be blank";
    } else if (!isEmail(formData.email)) {
      errors.email = "Please enter a valid email";
    }

    if (isEmpty(formData.password)) {
      errors.password = "Password can't be blank";
    } else if (isContainWhiteSpace(formData.password)) {
      errors.password = "Password should not contain white spaces";
    } else if (!isLength(formData.password, { gte: 5, lte: 15, trim: true })) {
      errors.password = "Password's length must between 5 to 15";
    }

    setFormState((prevState) => ({
      ...prevState,
      errors: errors,
    }));

    return errors;
  };

  const login = async (e) => {
    e.preventDefault();

    let errors = validateLoginForm();

    if (isEmpty(errors)) {
      const { formData } = formState;

      try {
        const resp = await axiosInstance.post("/login", {
          email: formData.email,
          password: formData.password,
        });
        const { body } = resp.data;
        setToken(body.token);
        navigate("/poll", { replace: true });
      } catch (err) {
        if (err.response.data) {
          errors.password = err.response.data.error;
        }
        setFormState((prevState) => ({
          ...prevState,
          errors: errors,
        }));
      }
    }
  };

  const handleInputChange = (event) => {
    const target = event.target;
    const value = target.value;
    const name = target.name;

    const updatedFormData = { ...formState.formData };
    updatedFormData[name] = value;
    setFormState((prevState) => ({
      ...prevState,
      formData: updatedFormData,
    }));
  };

  return (
    <div className="container-fluid d-flex justify-content-center align-items-center vh-100">
      <div className="card my-5" style={{ width: "500px" }}>
        <form className="card-body cardbody-color p-lg-5" onSubmit={login}>
          <div className="text-center">
            <img
              src="/logo.png"
              className="img-fluid profile-image-pic img-thumbnail rounded-circle my-3"
              width="200px"
              alt="profile"
            />
          </div>

          <div className="mb-3" id="email-row">
            <input
              type="email"
              className="form-control"
              id="email"
              aria-describedby="emailHelp"
              name="email"
              value={formState.formData?.email}
              placeholder="Email"
              onChange={handleInputChange}
            />
          </div>
          {formState.errors.email && (
            <div id="error-message-section">
              <p id="error-message" style={{ color: "red" }}>
                {formState.errors.email}
              </p>
            </div>
          )}

          <div className="mb-3" id="password-row">
            <input
              type="password"
              className="form-control"
              id="password"
              placeholder="Password"
              name="password"
              value={formState.formData?.password}
              onChange={handleInputChange}
            />
          </div>
          {formState.errors.password && (
            <div id="error-message-section">
              <p id="error-message" style={{ color: "red" }}>
                {formState.errors.password}
              </p>
            </div>
          )}

          <div className="text-center">
            <button
              type="submit"
              id="btn-submit"
              className="btn btn-color px-5 mb-5 w-100"
            >
              Login
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default Login;
