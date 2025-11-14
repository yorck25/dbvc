import {useEffect} from "react";
import {useParams} from "react-router-dom";
import {useDatabaseWorkerContext} from "../../contexts/databaseWorker.context.tsx";

export const DatabaseBrowser = () => {
    const {fetchDatabaseStructure} = useDatabaseWorkerContext();
    const {projectId} = useParams();

    useEffect(() => {
        fetchDatabaseStructure(Number(projectId)).then((structure) => {
            if(structure) {
                console.log(structure);
            } else {
                console.log("Failed to fetch database structure");
            }
        })
    }, [])

    return (
        <div></div>
    )
}