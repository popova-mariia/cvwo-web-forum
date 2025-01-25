import { useState } from "react";
import Nav from "./Nav";
import { makeStyles } from "@mui/styles";
import axios from "axios";

const useStyles = makeStyles(() => ({
    commentBody: {
        fontSize: 16,
        whiteSpace: "pre-wrap",
        paddingBottom: "1em",
    },
    commentCard: {
        marginBottom: "1em",
    },
    metadata: {
        fontSize: 14,
    },
}));

const Home = () => {
    const classes = useStyles();

    const [thread, setThread] = useState("");

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        try {
            const response = await axios.post('http://localhost:8080/api/threads', { title: thread });
            console.log('Thread created:', response.data);
            setThread("");
        } catch (error) {
            console.error('Error creating thread:', error);
        }
    };

    return (
        <>
            <Nav />
            <main className='home'>
                <h2 className='homeTitle'>Create a new Thread</h2>
                <form className='homeForm' onSubmit={handleSubmit}>
                    <div className='home__container'>
                        <label htmlFor='thread'>Title / Description</label>
                        <input
                            type='text'
                            name='thread'
                            required
                            value={thread}
                            onChange={(e) => setThread(e.target.value)}
                            className={classes.commentBody}
                        />
                    </div>
                    <button className='homeBtn'>CREATE THREAD</button>
                </form>
            </main>
        </>
    );
};

export default Home;