import logo from './mars.gif';
import './styles.css';

function App() {
  return (
    <div className='mars'>
          <img src={logo} className='App-logo' alt="mars" 
              height={700}
              width={700}
          />
      <h1 className='title-mars'>MARS</h1>
      <h1 className='sub-title'>THE RED PLANET</h1> 
      
    <div className='mars-weather'>
      <div className='date'>
        <h2 className='section-title section-title-date'>Sol
            <span data-current-sol> 375</span>
        </h2>
        <p className='date-day'>September 31</p>
      </div>

      <div className='temp'>
        <h2 className='section-title'>Temperature</h2>
        <p className='reading'>High: -20°C</p>
        <p className='reading'>Low: -20°C</p>
      </div>

      <div className='pressure'>
        <h2 className='section-title'>Pressure</h2>
        <p className='reading'>836 Pa -  <span data-pressure> [Higher] </span></p>
      </div>

      <div className='info'>
        <p>Curiosty is taking daily weather measurements (temperature, wind, and pressure) on the 
           surface of Mars at Gale Crater.
        </p> 
        <p>This is only a partial part of Curiosty's mission. 
            <a href='https://mars.nasa.gov/msl/home/'> Click here </a>
              to find out more.
        </p>
      </div>
    </div>
    </div> 
  );
}

export default App;
