import './app.css'
import { Sidebar } from './components/sidebar';
import {Router} from "./lib/router";

export function App() {
  return (
    <div>
        <Sidebar/>

        <Router/>
    </div>
  )
}
