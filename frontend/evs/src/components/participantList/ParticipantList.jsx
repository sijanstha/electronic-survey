import React, { useEffect, useState } from "react";
import { axiosInstance } from "../../axiosConfig";
import Sidebar from "../fragment/Sidebar";
import Navbar from "../fragment/Navbar";
import { DateTime } from "luxon";

const ParticipantList = () => {

    const [participantListFilter, setParticipantListFilter] = useState({
        sort: "desc",
        sortBy: "updated_at",
        limit: 5,
        page: 1,
        search: ""
    });

    const [apiResponse, setApiResponse] = useState({
        pageSize: 0,
        page: 0,
        totalRecords: 0,
        totalPages: 0,
        data: [],
    });

    // TODO: make an api call to list down valid sorting fields for poll
    const validSortFields = [
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
            title: "Name",
            value: "name",
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

        const updatedFilter = { ...participantListFilter };
        updatedFilter.page = page;
        setParticipantListFilter(updatedFilter);
    };

    const handleSortFieldChange = (e) => {
        e.preventDefault();
        const target = e.target;
        const updatedFilter = { ...participantListFilter };
        updatedFilter.sortBy = target.value;
        setParticipantListFilter(updatedFilter);
    };

    const handleSortDirectionChange = (e) => {
        e.preventDefault();
        const target = e.target;
        const updatedFilter = { ...participantListFilter };
        updatedFilter.sort = target.value;
        setParticipantListFilter(updatedFilter);
    };

    // TODO: make some delay while searching to reduce instant network call per user entered character
    const handleSearch = e => {
        const target = e.target;
        const updatedFilter = { ...participantListFilter };
        updatedFilter.search = target.value;
        setParticipantListFilter(updatedFilter);
    }
   
    useEffect(() => {
        const getPolls = async () => {
            try {
                const response = await axiosInstance.get(
                    `/participant-list?&sort=${participantListFilter.sort}&sortBy=${participantListFilter.sortBy}&limit=${participantListFilter.limit}&page=${participantListFilter.page}&search=${participantListFilter.search}`
                );
                const { data } = response;
                return data;
            } catch (err) {
                console.log(err);
            }
        };
        
        getPolls().then(
            result => setApiResponse(result));
    }, [participantListFilter]);

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
                                    <a className="btn btn-dark" href="/participant-list/add">
                                        Add Participant List
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
                                        onChange={handleSortFieldChange}
                                    >
                                        {validSortFields.length > 0 &&
                                            validSortFields.map((field) => (
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
                                        onChange={handleSortDirectionChange}
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
                                <div>
                                    <label>Search:</label>
                                    <input type="text" className="form-input form-input-sm" onChange={handleSearch}/>
                                </div>
                            </div>
                            <hr />
                            <table className="table">
                                <thead>
                                    <tr>
                                        <th scope="col">#</th>
                                        <th scope="col">Name</th>
                                        <th scope="col">Emails</th>
                                        <th scope="col">Created At</th>
                                        <th scope="col">Updated At</th>
                                        <th scope="col">Action</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {apiResponse?.data && (
                                        apiResponse.data.map((pl, idx) => (
                                            <tr key={pl.id}>
                                                <th scope="row">{idx + 1}</th>
                                                <td>{pl.name}</td>
                                                <td>{pl.emails.join(", ")}</td>
                                                <td>{DateTime.fromISO(pl.createdAt).toFormat('EEE, MMM yyyy hh:mm a')}</td>
                                                <td>{DateTime.fromISO(pl.updatedAt).toFormat('EEE, MMM yyyy hh:mm a')}</td>
                                                <td className="d-flex">
                                                    <div className="d-flex justify-content-center align-items-center gap-2">
                                                        <div>
                                                            <a href={`/participant-list/edit/${pl.id}`} title="Edit Poll">
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

export default ParticipantList;
