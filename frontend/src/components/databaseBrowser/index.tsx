import {useEffect} from "react";
import {NetworkAdapter, setAuthHeader} from "../../lib/networkAdapter.tsx";
import {useParams} from "react-router-dom";

export const DatabaseBrowser = () => {
    const {projectId} = useParams();

    useEffect(() => {
        console.log(projectId);

        const header = setAuthHeader();
        header.append("Content-Type", "application/json");

        const requestOptions: RequestInit = {
            method: NetworkAdapter.POST,
            headers: header,
        }

        fetch('http://localhost:8080/get-database-structure', requestOptions)
            .then(res => res.json())
            .then((data: any) => {
                console.log(data)
            });
    }, [])

    return (
        <div>

        </div>
    )
}