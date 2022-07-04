import React from 'react';
import {v4 as uuidv4} from 'uuid';

const createNewRoom = () => {
    let newRoomId = uuidv4();
    console.log(newRoomId);
}

const Home = () => {
    return (
        <div className="home">
            <button onClick={createNewRoom}>
                Create New Room
            </button>
        </div>
    )    
};

export default Home;
