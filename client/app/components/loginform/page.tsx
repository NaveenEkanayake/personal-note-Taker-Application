"use client";
import { motion } from "framer-motion";
import { useState, useEffect } from "react";
import Usericon from "../../../public/assests/user.webp";
import email from "../../../public/assests/email_icon.svg";
import lock from "../../../public/assests/lockicon.svg";
import crossicon from "@/public/assests/cross_icon.svg";
import Image from "next/image";
import LoginButton from "@/app/components/loginbutton/LoginButton";

type FormState = "Login" | "Signup";

export default function LoginForm() {
  const [state, setState] = useState<FormState>("Login");
  const [showLogin, setShowLogin] = useState<boolean>(true);

  useEffect(() => {
    document.body.style.overflow = "hidden";
    return () => {
      document.body.style.overflow = "unset";
    };
  }, []);
  const closeLoginForm = (): void => setShowLogin(false);
  if (!showLogin) return null;

  return (
    <div className="fixed top-0 left-0 right-0 bottom-0 z-10 backdrop-blur-sm bg-black/30 flex justify-center items-center">
      <motion.form
        initial={{ opacity: 0.2, y: 100 }}
        transition={{ duration: 1 }}
        animate={{ opacity: 1, y: 0 }}
        viewport={{ once: true }}
        className="relative bg-white p-10 rounded-xl text-slate-500 w-96"
      >
        <h1 className="text-center text-2xl text-neutral-700 font-medium">
          {state === "Login" ? "Login" : "Signup"}
        </h1>
        <p className="text-sm text-center">
          {state === "Login"
            ? "Welcome back! Please login to continue."
            : "Welcome! Please sign up to get started."}
        </p>

        {state !== "Login" && (
          <div className="border px-6 py-2 flex items-center gap-2 rounded-full mt-4">
            <Image src={Usericon} className="w-6 h-6" alt="User Icon" />
            <input
              type="text"
              name="fullname"
              placeholder="Fullname"
              required
              className="outline-none text-sm flex-grow"
            />
          </div>
        )}

        <div className="border px-6 py-2 flex items-center gap-2 rounded-full mt-4">
          <Image src={email} className="w-6 h-6" alt="Email Icon" />
          <input
            type="email"
            name="email"
            placeholder="Email"
            required
            className="outline-none text-sm flex-grow"
          />
        </div>

        <div className="border px-6 py-2 flex items-center gap-2 rounded-full mt-4">
          <Image src={lock} className="w-6 h-6" alt="Lock Icon" />
          <input
            type="password"
            name="password"
            placeholder="Password"
            required
            className="outline-none text-sm flex-grow"
          />
        </div>

        {state === "Login" && (
          <p className="text-sm text-blue-600 my-4 cursor-pointer text-center">
            Forgot Password?
          </p>
        )}
        <div className="flex justify-center mt-4">
          <LoginButton>
            {state === "Login" ? "Login" : "Create Account"}
          </LoginButton>
        </div>

        <p className="mt-5 text-center">
          {state === "Login"
            ? "Don't have an account?"
            : "Already have an account?"}{" "}
          <span
            className="text-blue-600 cursor-pointer"
            onClick={() => setState(state === "Login" ? "Signup" : "Login")}
          >
            {state === "Login" ? "Sign Up" : "Login"}
          </span>
        </p>
        <Image
          src={crossicon}
          className="absolute top-5 right-5 cursor-pointer"
          alt="Close"
          onClick={closeLoginForm}
        />
      </motion.form>
    </div>
  );
}
