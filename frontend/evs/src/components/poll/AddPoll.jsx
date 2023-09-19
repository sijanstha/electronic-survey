import React, { useState } from "react";
import Sidebar from "../fragment/Sidebar";
import Navbar from "../fragment/Navbar";
import { isEmpty } from "../../shared/validator";
import { axiosInstance } from "../../axiosConfig";
import { useNavigate } from "react-router-dom";
import {DateTime} from "luxon";
import { useAlert } from "react-alert";

const AddPoll = () => {
    const navigate = useNavigate();
    const alert = useAlert();

    const [formState, setFormState] = useState({
        formData: { title: '', description: '', strStartsAt: '', strEndsAt: '', timezone: '', startsAt:'', endsAt:'' },
        errors: { title: '', startsAt: '', endsAt: '', timezone: '' }
    });

    const validateAddPollForm = () => {
        let errors = {};
        const { formData } = formState;

        if (isEmpty(formData.title)) {
            errors.title = "Title can't be blank";
        }

        if (isEmpty(formData.strStartsAt)) {
            errors.startsAt = "Start date can't be blank";
        }

        if (isEmpty(formData.strEndsAt)) {
            errors.endsAt = "End date can't be blank";
        }

        if (isEmpty(formData.timezone)) {
            errors.timezone = "Timezone can't be blank";
        }

        const startsAt = DateTime.fromISO(formData.strStartsAt, {zone: formData.timezone}).toUTC();
        const endsAt = DateTime.fromISO(formData.strEndsAt, {zone: formData.timezone}).toUTC();
        const today = DateTime.local().toUTC();

        if (startsAt.diff(today) <= 0) {
            errors.startsAt = "Start date can't be of past date";
        }

        if (endsAt.diff(today) <= 0) {
            errors.endsAt = "End date can't be of past date";
        }

        if (endsAt.diff(startsAt) <= 0) {
            errors.endsAt = "End date should be after start date";
        }

        formData.startsAt = startsAt.toISO();
        formData.endsAt = endsAt.toISO();

        setFormState((prevState) => ({
            ...prevState,
            errors: errors,
        }));

        return errors;
    };

    const handleAddPoll = async (e) => {
        e.preventDefault();

        let errors = validateAddPollForm();
        if (isEmpty(errors)) {
            const { formData } = formState;
            try {
                await axiosInstance.post("/poll", {
                    ...formData
                });
                alert.success("Poll added successfully");
                navigate("/poll", { replace: true });
            } catch (err) {
                console.log('errss', err)
                if (err.response.data) {
                    errors.title = err.response.data.error;
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
        <div id="viewport">
            <Sidebar />
            <Navbar />
            <div className="container">
                <div id="content">
                    <div className="container-fluid">
                        <h4>Add Poll</h4>
                        <hr />
                        <form className="card-body cardbody-color p-lg-5" onSubmit={handleAddPoll}>
                            <div className="form-group mb-3 row">
                                <div className="col-1">
                                    <label htmlFor="title" className="col-form-label">Title</label>
                                </div>
                                <div className="col-11">
                                    <div>
                                        <input type="text"
                                            className="form-control w-50"
                                            id="title"
                                            name="title"
                                            value={formState.formData?.title}
                                            onChange={handleInputChange}
                                        />
                                    </div>
                                    {formState.errors.title && (
                                        <div id="error-message-section">
                                            <p id="error-message" style={{ color: "red" }}>
                                                {formState.errors.title}
                                            </p>
                                        </div>
                                    )}
                                </div>
                            </div>
                            <div className="form-group mb-3 row">
                                <div className="col-1">
                                    <label htmlFor="description" className="col-form-label">Description</label>
                                </div>
                                <div className="col-11">
                                    <div>
                                        <textarea
                                            className="form-control w-50"
                                            id="description"
                                            name="description"
                                            value={formState.formData?.description}
                                            onChange={handleInputChange}
                                        />
                                    </div>
                                </div>
                            </div>
                            <div className="form-group mb-3 row">
                                <div className="col-1">
                                    <label htmlFor="startsAt" className="col-form-label">Starts At</label>
                                </div>
                                <div className="col-11">
                                    <input
                                        className="form-control w-50"
                                        type="datetime-local"
                                        id="strStartsAt"
                                        name="strStartsAt"
                                        value={formState.formData?.strStartsAt}
                                        onChange={handleInputChange}
                                    />
                                    {formState.errors.startsAt && (
                                        <div id="error-message-section">
                                            <p id="error-message" style={{ color: "red" }}>
                                                {formState.errors.startsAt}
                                            </p>
                                        </div>
                                    )}
                                </div>
                            </div>

                            <div className="form-group mb-3 row">
                                <div className="col-1">
                                    <label htmlFor="endsAt" className="col-form-label">Ends At</label>
                                </div>
                                <div className="col-11">
                                    <input
                                        className="form-control w-50"
                                        type="datetime-local"
                                        id="strEndsAt"
                                        name="strEndsAt"
                                        value={formState.formData?.strEndsAt}
                                        onChange={handleInputChange}
                                    />
                                    {formState.errors.endsAt && (
                                        <div id="error-message-section">
                                            <p id="error-message" style={{ color: "red" }}>
                                                {formState.errors.endsAt}
                                            </p>
                                        </div>
                                    )}
                                </div>
                            </div>

                            <div className="form-group mb-3 row">
                                <div className="col-1">
                                    <label htmlFor="timezone" className="col-form-label">Timezone</label>
                                </div>
                                <div className="col-11">
                                    <select
                                        className="form-control w-50"
                                        aria-label="Timezone"
                                        onChange={handleInputChange}
                                        id="timezone"
                                        name="timezone"
                                    >
                                        {
                                            Intl.supportedValuesOf('timeZone').map(tz => <option key={tz} value={tz}>{tz}</option>)
                                        }
                                    </select>
                                    {formState.errors.timezone && (
                                        <div id="error-message-section">
                                            <p id="error-message" style={{ color: "red" }}>
                                                {formState.errors.timezone}
                                            </p>
                                        </div>
                                    )}
                                </div>
                            </div>

                            <div className="row">
                                <div className="col-1">
                                </div>
                                <div className="col-11">
                                    <div className="d-flex justify-content-end w-50">
                                        <button type="submit" className="btn btn-sm btn-dark mt-2">Add Poll</button>
                                    </div>
                                </div>
                            </div>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default AddPoll;