import React, { useEffect, useState } from "react";
import { axiosInstance } from "../../axiosConfig";


const Poll = () => {
    const [apiResponse, setApiResponse] = useState({
        pageSize: 0,
        page: 0,
        totalRecords: 0,
        data: []
    });
    const [pollListFilter, setPostListFilter] = useState({
        sort: 'desc',
        sortBy: 'updated_at',
        limit: 2,
        page: 1,
        showOwnPoll: true,
    });

    const getPolls = async () => {
        try {
            const response = await axiosInstance.get(`/poll?showOwnPoll=${pollListFilter.showOwnPoll}&sort=${pollListFilter.sort}&sortBy=${pollListFilter.sortBy}&limit=${pollListFilter.limit}&page=${pollListFilter.page}`);
            const { data } = response;
            return data;
        } catch (err) {
            console.log(err)
        }
    };

    useEffect(() => {
        getPolls().then(
            result => setApiResponse(result));
    }, []);

    return (
        <React.Fragment>
            <div id="viewport">
                <div id="sidebar">
                    <header>
                        <a href="/">Electronic Voting System</a>
                    </header>
                    <ul class="nav">
                        <li>
                            <a href="/">
                                <i class="zmdi zmdi-view-dashboard"></i> Dashboard
                            </a>
                        </li>
                        <li>
                            <a href="/admin/organizers">
                                <i class="zmdi zmdi-link"></i> Manage Organizers
                            </a>
                        </li>
                        <li>
                            <a href="/poll">
                                <i class="zmdi zmdi-calendar"></i> Manage Polls
                            </a>
                        </li>
                    </ul>
                </div>
                <nav class="navbar navbar-expand-lg bg-body-tertiary mb-5 p-2" style={{ "box-shadow": "rgba(0, 0, 0, 0.1) 0px 4px 12px;" }}>
                    <div class="container-fluid">
                        <img class="navbar-brand" src="/logo.png" width="50px" alt="" />
                        <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
                            <span class="navbar-toggler-icon"></span>
                        </button>
                        <div class="collapse navbar-collapse" id="navbarSupportedContent">
                            <ul class="navbar-nav me-auto mb-2 mb-lg-0">

                            </ul>
                            <div class="d-flex list-style-none" style={{ "list-style": "none" }}>
                                <li class="nav-item dropdown">
                                    <a class="nav-link dropdown-toggle" href="" role="button" data-bs-toggle="dropdown" aria-expanded="false">
                                        Email
                                    </a>
                                    <ul class="dropdown-menu">
                                        <li><a class="dropdown-item" href="/logout">Logout</a></li>
                                    </ul>
                                </li>
                            </div>
                        </div>
                    </div>
                </nav>
                <div class="container">
                    <div id="content">
                        <div class="container-fluid">
                            <div class="d-flex justify-content-between">
                                <h4>Polls</h4>
                                <div>
                                    <a class="btn btn-dark" href="/poll/add">Add Poll</a>
                                </div>
                            </div>

                            <table class="table">
                                <thead>
                                    <tr>
                                        <th scope="col">#</th>
                                        <th scope="col">Title</th>
                                        <th scope="col">Description</th>
                                        <th scope="col">Starts At</th>
                                        <th scope="col">Ends At</th>
                                        <th scope="col">State</th>
                                        <th scope="col">Primary Organizer</th>
                                        <th scope="col">Action</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {apiResponse.data && (
                                        apiResponse.data.map((poll, idx) => (
                                            <tr key={poll.id}>
                                                <th scope="row">{idx + 1}</th>
                                                <td>{poll.title}</td>
                                                <td>{poll.description}</td>
                                                <td>{poll.startsAt}</td>
                                                <td>{poll.endsAt}</td>
                                                <td>{poll.state}</td>
                                                <td>{poll.primaryOrganizerName}</td>
                                                <td class="d-flex">
                                                    <div class="d-flex justify-content-center align-items-center gap-2">
                                                        <div>
                                                            <a href="#"><i class="fa fa-trash"></i></a>
                                                        </div>
                                                        <div>
                                                        <a href="#"><i class="fa fa-edit"></i></a>
                                                        </div>
                                                    </div>
                                                </td>
                                            </tr>
                                        ))
                                    )
                                    }
                                </tbody>
                            </table>
                        </div>
                    </div>
                </div>
            </div>
        </React.Fragment>
    );
}

export default Poll;