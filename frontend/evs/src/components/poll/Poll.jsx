import React, { useEffect, useState } from "react";
import { axiosInstance } from "../../axiosConfig";
import Sidebar from "../fragment/Sidebar";
import Navbar from "../fragment/Navbar";
import { DateTime } from "luxon";

const Poll = () => {
    const validPollStates = [
        { isChecked: true, value: "PREPARED" },
        { isChecked: true, value: "STARTED" },
        { isChecked: true, value: "VOTING" },
        { isChecked: true, value: "FINISHED" },
    ];

    const [pollListFilter, setPostListFilter] = useState({
        sort: "desc",
        sortBy: "updated_at",
        limit: 5,
        page: 1,
        showOwnPoll: true,
        states: validPollStates,
    });

    const [apiResponse, setApiResponse] = useState({
        pageSize: 0,
        page: 0,
        totalRecords: 0,
        totalPages: 0,
        data: [],
    });

    const [showDropdown, setShowDropdown] = useState(false);

    // TODO: make an api call to list down valid sorting fields for poll
    const validPollSortFields = [
        {
            title: "Updated At",
            value: "updated_at",
            isDefaultSort: true,
        },
        {
            title: "Created At",
            value: "created_at",
            isDefaultSort: false,
        },
        {
            title: "Id",
            value: "id",
            isDefaultSort: false,
        },
        {
            title: "State",
            value: "state",
            isDefaultSort: false,
        },
        {
            title: "Title",
            value: "title",
            isDefaultSort: false,
        },
    ];
    const validSortDirection = [
        {
            title: "ASC",
            value: "asc",
            isDefaultSort: false,
        },
        {
            title: "DESC",
            value: "desc",
            isDefaultSort: true,
        },
    ];

    const handlePagination = (e) => {
        e.preventDefault();
        const target = e.target;
        const page = parseInt(target.getAttribute("page"));

        const updatedFilter = { ...pollListFilter };
        updatedFilter.page = page;
        setPostListFilter(updatedFilter);
    };

    const handlePollSortFieldChange = (e) => {
        e.preventDefault();
        const target = e.target;
        const updatedFilter = { ...pollListFilter };
        updatedFilter.sortBy = target.value;
        setPostListFilter(updatedFilter);
    };

    const handlePollSortDirectionChange = (e) => {
        e.preventDefault();
        const target = e.target;
        const updatedFilter = { ...pollListFilter };
        updatedFilter.sort = target.value;
        setPostListFilter(updatedFilter);
    };

    let pollValueArray = [];
    const handlePollStateChange = (e, i) => {
        const updatedStates = [...pollListFilter.states];

        if (e.target.checked) {
            pollValueArray.push(e.target.value);

            updatedStates[i] = {
                ...updatedStates[i],
                isChecked: true,
            };

        } else {
            const index = pollValueArray.findIndex(
                (value) => value === e.target.value
            );
            pollValueArray.splice(index, 1);

            updatedStates[i] = {
                ...updatedStates[i],
                isChecked: false,
            };
        }

        setPostListFilter((prevState) => ({
            ...prevState,
            page: 1,
            states: updatedStates,
        }));
    };

    useEffect(() => {
        const getPolls = async () => {
            try {
                let state = "(";
                state += pollListFilter.states.filter(state => state.isChecked).map(state => state.value).join(",").concat(")");
                const response = await axiosInstance.get(
                    `/poll?showOwnPoll=${pollListFilter.showOwnPoll}&sort=${pollListFilter.sort}&sortBy=${pollListFilter.sortBy}&limit=${pollListFilter.limit}&page=${pollListFilter.page}&state=${state}`
                );
                const { data } = response;
                return data;
            } catch (err) {
                console.log(err);
            }
        };
        
        getPolls().then(
            result => setApiResponse(result));
    }, [pollListFilter]);

    return (
        <React.Fragment>
            <div id="viewport">
                <Sidebar />
                <Navbar />

                <div className="container">
                    <div id="content">
                        <div className="container-fluid">
                            <div className="d-flex justify-content-between">
                                <h4>Polls</h4>
                                <div>
                                    <a className="btn btn-dark" href="/poll/add">
                                        Add Poll
                                    </a>
                                </div>
                            </div>
                            <hr />
                            <div className="d-flex justify-content-between">
                                <div>
                                    <label>Sort By:</label>
                                    <select
                                        className="form-select form-select-sm"
                                        aria-label="Sort by"
                                        onChange={handlePollSortFieldChange}
                                    >
                                        {validPollSortFields.length > 0 &&
                                            validPollSortFields.map((field) => (
                                                <option
                                                    key={field.value}
                                                    selected={field.isDefaultSort}
                                                    value={field.value}
                                                >
                                                    {field.title}
                                                </option>
                                            ))}
                                    </select>
                                </div>
                                <div>
                                    <label>Sort Direction:</label>
                                    <select
                                        className="form-select form-select-sm"
                                        aria-label="Sort direction"
                                        onChange={handlePollSortDirectionChange}
                                    >
                                        {validSortDirection.length > 0 &&
                                            validSortDirection.map((field) => (
                                                <option
                                                    key={field.value}
                                                    selected={field.isDefaultSort}
                                                    value={field.value}
                                                >
                                                    {field.title}
                                                </option>
                                            ))}
                                    </select>
                                </div>
                                <div className="position-relative">
                                    <label onClick={() => setShowDropdown(!showDropdown)}>Filter by Poll state:</label>
                                    {showDropdown && <div className="filter-dropdown shadow">
                                        {pollListFilter?.states.map((poll, index) => {
                                            return (
                                                <div
                                                    className="d-flex align-items-baseline gap-2"
                                                    key={index}
                                                >
                                                    <input
                                                        type="checkbox"
                                                        checked={poll.isChecked}
                                                        name={`${index}`}
                                                        value={poll.value}
                                                        onChange={(e) => handlePollStateChange(e, index)}
                                                    />
                                                    <div>{poll.value}</div>
                                                </div>
                                            );
                                        })}
                                    </div>}

                                </div>
                            </div>
                            <hr />
                            <table className="table">
                                <thead>
                                    <tr>
                                        <th scope="col">#</th>
                                        <th scope="col">Title</th>
                                        <th scope="col">Description</th>
                                        <th scope="col">Starts At</th>
                                        <th scope="col">Ends At</th>
                                        <th scope="col">Timezone</th>
                                        <th scope="col">State</th>
                                        <th scope="col">Primary Organizer</th>
                                        <th scope="col">Action</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {apiResponse?.data && (
                                        apiResponse.data.map((poll, idx) => (
                                            <tr key={poll.id}>
                                                <th scope="row">{idx + 1}</th>
                                                <td>{poll.title}</td>
                                                <td>{poll.description.length > 50 ?
                                                    `${poll.description.substring(0, 50)}...` : poll.description
                                                }</td>
                                                <td>{DateTime.fromISO(poll.startsAt).setZone(poll.timezone).toFormat('EEE, MMM yyyy hh:mm a')}</td>
                                                <td>{DateTime.fromISO(poll.endsAt).setZone(poll.timezone).toFormat('EEE, MMM yyyy hh:mm a')}</td>
                                                <td>{poll.timezone}</td>
                                                <td>{poll.state}</td>
                                                <td>{poll.primaryOrganizerName}</td>
                                                <td className="d-flex">
                                                    <div className="d-flex justify-content-center align-items-center gap-2">
                                                        <div title="Start Poll">
                                                            <i className="fa fa-play"></i>
                                                        </div>
                                                        <div>
                                                            <a href={`/poll/edit/${poll.id}`} title="Edit Poll">
                                                                <i className="fa fa-edit"></i>
                                                            </a>
                                                        </div>
                                                        <div title="Delete Poll">
                                                            <i className="fa fa-trash"></i>
                                                        </div>
                                                    </div>
                                                </td>
                                            </tr>
                                        )))}
                                </tbody>
                            </table>
                            <div className="d-flex ml-4">
                                <nav aria-label="Page navigation example">
                                    <ul className="pagination">
                                        {apiResponse?.data && ([...Array(apiResponse.totalPages)].map((x, i) =>
                                            <li className="page-item" key={i}>
                                                <a className="page-link" page={i + 1} href=" " onClick={handlePagination}>{i + 1}</a>
                                            </li>
                                        ))}
                                    </ul>
                                </nav>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </React.Fragment>
    );
};

export default Poll;
