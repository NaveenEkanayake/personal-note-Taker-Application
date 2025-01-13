"use client";
import { useState } from "react";
import Description from "./components/description/Description";
import Footer from "./components/footer/Footer";
import NavBar from "./components/navbar/NavBar";
import LoginForm from "./components/loginform/page";
export default function Home() {
  const [showLogin, setShowLogin] = useState(false);
  return (
    <>
      {showLogin && <LoginForm />}
      <NavBar setShowLogin={setShowLogin} />
      <Description />
      <Footer />
    </>
  );
}
