import React from 'react';
import './Benefits3.css';
import LG from './assets/LG.png'
import pointsicon from './assets/pointsicon.png'
import { useNavigate } from "react-router-dom"

const Benefits3 = () => {
  const navigate = useNavigate();

    const handleBenefitsClick = () => {
      navigate("/");
    };

    const handleBenefits2Click = () => {
      navigate("/Benefits2");
    };

    const handleBenefits3Click = () => {
      navigate("/Benefits3");
    };

  return (
    <div className="airline-selection">
      <header className="header">
        {/* <div className="logo"></div>
        <div className="app-name">MawMaw</div> */}
        <img className="head-logo" src={LG} alt="logo" />
        <div className="navbar1">
          <button onClick={handleBenefitsClick}>Home</button>
          {/* <span className='home'>Home</span> */}
        </div>
        <div className="navbar2">
          <button onClick={handleBenefits2Click}>Flight</button>
          {/* <span className='flight'>Flight</span> */}
        </div>
        <div className="navbar3">
          <button onClick={handleBenefits3Click}>Benefits</button>
          {/* <span className='benefits'>Benefits</span> */}
        </div>
        <div className="navbar4">
          <button>Help Center</button>
          {/* <span className='help-center'>Help Center</span> */}
        </div>
        {/* <div className="search-bar"></div> */}
      </header>

      <div className="frame-flight-details"></div>
      <div className="flight-details">
        <div className="header2"></div>
        <div className="flight-info">
          <span className='ff-gt'>Bangkok(Suvarnabhumi) - Samui</span>
            {/* <h2>Bangkok(Suvarnabhumi) - Samui</h2>
            <p>Airline 1</p> */}
          <span className='airline-name'>Airline 1</span>
          <div className="a-points">
            <img className="p-icon" src={pointsicon} alt="icon" />
            <span className="points">6,500 P</span>
          </div>

          <span className='class'>Class :</span>
          <div className="class-class">
            <div className="eco-button">
              <button>Economy</button>
            </div>
            <div className="bu-button">
              <button>Business</button>
            </div>
          </div>

          <span className='trip'>Trip :</span>
          <div className="trip-class">
            <div className="ow-button">
              <button>One Way</button>
            </div>
            <div className="rt-button">
              <button>Round Trip</button>
            </div>
          </div>
          {/* <div className="trip">
            <label>Trip :</label>
            <button>One Way</button>
            <button>Round Trip</button>
          </div> */}
          <button className="redeem-button">Redeem benefits</button>
        </div>
      </div>
    </div>
  );
};

export default Benefits3;
