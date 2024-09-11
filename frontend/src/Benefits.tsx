import React from 'react';
import './Benefits.css';
import LG from './assets/LG.png'
import { useNavigate } from "react-router-dom"

export default function Benefits() {
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
    <div className='main-container'>
      <div className='flex-row'>
        <div className='rectangle'>
          <img className="logo" src={LG} alt="logo" />

          {/* <div className='flex-row-dd'> */}
            <div className="rectangle-1">
              <button onClick={handleBenefitsClick}>Home</button>
              {/* <span className='home'>Home</span> */}
            </div>
            <div className="rectangle-2">
              <button onClick={handleBenefits2Click}>Flight</button>
              {/* <span className='flight'>Flight</span> */}
            </div>
            <div className="rectangle-3">
              <button onClick={handleBenefits3Click}>Benefits</button>
              {/* <span className='benefits'>Benefits</span> */}
            </div>
            <div className="rectangle-4">
              <button>Help Center</button>
              {/* <span className='help-center'>Help Center</span> */}
            {/* </div> */}
            {/* <div className='ellipse' /> */}
            {/* <div className='rectangle-1'>
              <span className='home'>Home</span>
            </div>
            <div className='rectangle-2'>
              <span className='flight'>Flight</span>
            </div>
            <div className='rectangle-3'>
              <span className='benefits'>Benefits</span>
            </div>
            <div className='rectangle-4'>
              <span className='help-center'>Help Center</span>
            </div> */}
            {/* <div className='rectangle-5' /> */}
          </div>
          {/* <div className='rectangle-6'>
            <span className='mawmaw'>MawMaw</span>
          </div> */}
        </div>
        <div className='rectangle-7'>
          <span className='select-airline'>Please select an airline</span>
        </div>
      </div>
      <div className='flex-row-ff'>
        <div className='rectangle-8' />
        <div className='rectangle-9' />
        <div className='rectangle-a' />
        <div className='rectangle-b' />
      </div>
      <div className='flex-row-fb'>
        <span className='airline-1'>Airline 1</span>
        <span className='airline-2'>Airline 2</span>
        <span className='airline-3'>Airline 3</span>
        <span className='airline-4'>Airline 4</span>
      </div>
    </div>
  );
}