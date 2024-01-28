import './App.css'
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import {Home} from "./routes/Home.tsx";
import {SignIn} from "./routes/SignIn.tsx";
import {AuthRoute} from "./routes/AuthRoute.tsx";
import {Reading} from "@/routes/Reading";


function App() {

  return (
      <BrowserRouter>
          <Routes>
              <Route path="/signin" element={<SignIn />} />
              <Route path="/"
                     element={<AuthRoute>
                         <Home/>
                     </AuthRoute>} />
              <Route path="/psalms/:id"
                     element={<AuthRoute>
                         <Reading/>
                     </AuthRoute>} />
          </Routes>
      </BrowserRouter>
  )
}

export default App
