import React, { Component } from "react";
import './login.css';
import { isContainWhiteSpace, isEmail, isEmpty, isLength } from "../../shared/validator";
import axios from 'axios'

class Login extends Component {

    constructor(props) {
        super(props)

        this.state = {
            formData: {}, // Contains login form data
            errors: {}, // Contains login field errors
            formSubmitted: false, // Indicates submit status of login form
            loading: false // Indicates in progress state of login form
        }
    }

    handleInputChange = (event) => {
        const target = event.target;
        const value = target.value;
        const name = target.name;

        let { formData } = this.state;
        formData[name] = value;

        this.setState({
            formData: formData
        });
    }

    validateLoginForm = (e) => {

        let errors = {};
        const { formData } = this.state;

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

        if (isEmpty(errors)) {
            return true;
        } else {
            return errors;
        }
    }

    login = (e) => {
        e.preventDefault();
        let errors = this.validateLoginForm();
        if (errors === true) {
            const { formData } = this.state;
            axios.post('http://localhost:4000/api/login' , {
                email: formData.email,
                password: formData.password
            }).then(resp => {
                alert("You are successfully signed in...");
                console.log(resp);
                window.location.reload()
            }).catch(err => {
                console.log(err);
            });
            
        } else {
            this.setState({
                errors: errors,
                formSubmitted: true
            });
        }
    }

    render() {

        const { errors, formSubmitted } = this.state;

        return (
            <div class="container-fluid d-flex justify-content-center align-items-center vh-100">
                <div class="card my-5" style={{ width: '500px' }}>
                    <form class="card-body cardbody-color p-lg-5" onSubmit={this.login}>

                        <div class="text-center">
                            <img src="/logo.png"
                                class="img-fluid profile-image-pic img-thumbnail rounded-circle my-3"
                                width="200px" alt="profile" />

                        </div>

                        <div class="mb-3" id="email-row" >
                            <input type="email" class="form-control" id="email" aria-describedby="emailHelp" name="email"
                                placeholder="Email" onChange={this.handleInputChange} />
                        </div>
                        {errors.email &&
                            <div id="error-message-section">
                                <p id="error-message" style={{ color: 'red' }}>{errors.email}</p>
                            </div>
                        }

                        <div class="mb-3" id="password-row">
                            <input type="password" class="form-control" id="password" placeholder="Password" 
                            name="password" onChange={this.handleInputChange} />
                        </div>
                        {errors.password &&
                            <div id="error-message-section">
                                <p id="error-message" style={{ color: 'red' }}>{errors.password}</p>
                            </div>
                        }

                        <div class="text-center">
                            <button type="submit" id="btn-submit" class="btn btn-color px-5 mb-5 w-100">Login</button>
                        </div>
                    </form>

                </div>
            </div>
        )
    }
}

export default Login;