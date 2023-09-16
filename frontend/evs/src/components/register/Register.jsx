import { useState } from "react";
import { isContainWhiteSpace, isEmail, isEmpty, isLength } from "../../shared/validator";
import { axiosInstance } from "../../axiosConfig";
import { useNavigate } from "react-router-dom";

const Register = () => {
    const navigate = useNavigate();
    const [formState, setFormState] = useState({
        formData: { firstName: '', lastName: '', email: '', password: '', repeatPassword: '' },
        errors: { firstName: '', lastName: '', email: '', password: '' }
    });

    const validateRegisterForm = () => {
        let errors = {};
        const { formData } = formState;

        if (isEmpty(formData.firstName)) {
            errors.firstName = "Firstname can't be blank";
        }

        if (isEmpty(formData.lastName)) {
            errors.lastName = "Lastname can't be blank";
        }

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

        if (isEmpty(formData.repeatPassword)) {
            errors.repeatPassword = "Password can't be blank";
        } else if (isContainWhiteSpace(formData.password)) {
            errors.repeatPassword = "Password should not contain white spaces";
        } else if (!isLength(formData.password, { gte: 5, lte: 15, trim: true })) {
            errors.repeatPassword = "Password's length must between 5 to 15";
        }

        if (formData.password !== formData.repeatPassword) {
            errors.repeatPassword = "Password doesn't match";
        }

        setFormState((prevState) => ({
            ...prevState,
            errors: errors,
        }));

        return errors;
    };

    const register = async (e) => {
        e.preventDefault();
        let errors = validateRegisterForm();
        console.log('err', errors)

        if (isEmpty(errors)) {
            const { formData } = formState;

            try {
                await axiosInstance.post("/register", {
                    firstName: formData.firstName,
                    lastName: formData.lastName,
                    email: formData.email,
                    password: formData.password,
                });
                navigate("/login", { replace: true });
            } catch (err) {
                console.log('errss', err)
                if (err.response.data) {
                    errors.email = err.response.data.error;
                }
                setFormState((prevState) => ({
                    ...prevState,
                    errors: errors,
                }));
            }
        }
    }

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
                <form className="card-body cardbody-color p-lg-5" onSubmit={register}>
                    <div className="text-center">
                        <img
                            src="/logo.png"
                            className="img-fluid profile-image-pic img-thumbnail rounded-circle my-3"
                            width="200px"
                            alt="profile"
                        />
                    </div>

                    <div className="mb-3" id="firstName-row">
                        <input
                            type="text"
                            className="form-control"
                            id="firstName"
                            name="firstName"
                            placeholder="First Name"
                            value={formState.formData?.firstName}
                            onChange={handleInputChange}
                        />
                    </div>
                    {formState.errors.firstName && (
                        <div id="error-message-section">
                            <p id="error-message" style={{ color: "red" }}>
                                {formState.errors.firstName}
                            </p>
                        </div>
                    )}

                    <div className="mb-3" id="lastName-row">
                        <input
                            type="text"
                            className="form-control"
                            id="lastName"
                            name="lastName"
                            placeholder="Last Name"
                            value={formState.formData?.lastName}
                            onChange={handleInputChange}
                        />
                    </div>
                    {formState.errors.lastName && (
                        <div id="error-message-section">
                            <p id="error-message" style={{ color: "red" }}>
                                {formState.errors.lastName}
                            </p>
                        </div>
                    )}

                    <div className="mb-3" id="email-row">
                        <input
                            type="email"
                            className="form-control"
                            id="email"
                            name="email"
                            placeholder="Email"
                            value={formState.formData?.email}
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

                    <div className="mb-3" id="repeatPassword-row">
                        <input
                            type="password"
                            className="form-control"
                            id="repeatPassword"
                            placeholder="Re-enter password"
                            name="repeatPassword"
                            value={formState.formData?.repeatPassword}
                            onChange={handleInputChange}
                        />
                    </div>
                    {formState.errors.repeatPassword && (
                        <div id="error-message-section">
                            <p id="error-message" style={{ color: "red" }}>
                                {formState.errors.repeatPassword}
                            </p>
                        </div>
                    )}

                    <div className="text-center">
                        <button
                            type="submit"
                            id="btn-submit"
                            className="btn btn-color px-5 mb-5 w-100"
                        >
                            Register
                        </button>
                    </div>
                    <div className="text-center">
                        Already have an account? Please <a href="/login">login</a>
                    </div>
                </form>
            </div>
        </div>
    );
}

export default Register;