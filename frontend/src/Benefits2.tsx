import React from 'react';
import './Benefits2.css';
import LG from './assets/LG.png'
import pointsicon from './assets/pointsicon.png'
import { useNavigate } from "react-router-dom"

const Benefits2: React.FC = () => {
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
    <div className="container">
      <div className="navbar">
        <img className="header-logo" src={LG} alt="logo" />
        {/* <div className="navbar-logo"></div> */}
          {/* <div className="circle-logo" /> */}
          {/* <div className="navbar-item mawmaw">MawMaw</div> */}
        <div className="navbar-links1">
          <button onClick={handleBenefitsClick}>Home</button>
          {/* <span className='home'>Home</span> */}
        </div>
        <div className="navbar-links2">
          <button onClick={handleBenefits2Click}>Flight</button>
          {/* <span className='flight'>Flight</span> */}
        </div>
        <div className="navbar-links3">
          <button onClick={handleBenefits3Click}>Benefits</button>
          {/* <span className='benefits'>Benefits</span> */}
        </div>
        <div className="navbar-links4">
          <button>Help Center</button>
          {/* <span className='help-center'>Help Center</span> */}
        </div>
      </div>

      <div className="airline-tabs">
        <div className="tab">Airline 1</div>
        <div className="tab">Airline 2</div>
        <div className="tab">Airline 3</div>
        <div className="tab">Airline 4</div>
      </div>

      <div className="flight-container1">
        <div className="flight-card-frame1"></div>
        <div className="flight-image1"></div>
        <div className="flight-info-frame1"></div>
        <div className="flight-info1">
          <div className="flight-name1">Bangkok (Suvarnabhumi) - Samui</div>
          <div className="flight-price1">
            <img className="p-icon1" src={pointsicon} alt="icon" />
            <div className="price1">6,500 P</div>
          </div>
        </div>
      </div>
  
      <div className="flight-container2">
        <div className="flight-card-frame2"></div>
        <div className="flight-image2"></div>
        <div className="flight-info-frame2"></div>
        <div className="flight-info2">
          <div className="flight-name2">Bangkok (Suvarnabhumi) - Maldives</div>
          <div className="flight-price2">
            <img className="p-icon2" src={pointsicon} alt="icon" />
            <div className="price2">14,000 P</div>
          </div>
        </div>
      </div>

      {/* <div className="flight-container3">
        <div className="flight-card-frame3"></div>
        <div className="flight-image3"></div>
        <div className="flight-info-frame3"></div>
        <div className="flight-info3">
          <div className="flight-name3">Bangkok (Suvarnabhumi) - Lampang</div>
          <div className="flight-price3">
            <img className="p-icon3" src={pointsicon} alt="icon" />
            <div className="price3">5,000 P</div>
          </div>
        </div>
      </div> */}

      <div className="flight-container4">
        <div className="flight-card-frame4"></div>
        <div className="flight-image4"></div>
        <div className="flight-info-frame4"></div>
        <div className="flight-info4">
          <div className="flight-name4">Bangkok (Suvarnabhumi) - Trat</div>
          <div className="flight-price4">
            <img className="p-icon4" src={pointsicon} alt="icon" />
            <div className="price4">5,000 P</div>
          </div>
        </div>
      </div>

      <div className="flight-container5">
        <div className="flight-card-frame5"></div>
        <div className="flight-image5"></div>
        <div className="flight-info-frame5"></div>
        <div className="flight-info5">
          <div className="flight-name5">Bangkok (Suvarnabhumi) - Phuket</div>
          <div className="flight-price5">
            <img className="p-icon5" src={pointsicon} alt="icon" />
            <div className="price5">5,000 P</div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Benefits2;
