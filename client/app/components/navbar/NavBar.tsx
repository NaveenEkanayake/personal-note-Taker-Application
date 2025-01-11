"use client";
import Link from "next/link";
import logo from "../../../public/assests/personalnotelogo.png";
import Image from "next/image";
import { Popover, PopoverButton, PopoverPanel } from "@headlessui/react";
import { Bars3Icon, XMarkIcon } from "@heroicons/react/24/outline";

export default function NavBar() {
  return (
    <Popover className="bg-transparent border-b-2 border-gray-300 h-24 px-6 flex items-center justify-between">
      <Link href="/">
        <div className="flex items-center">
          <h1 className="flex items-center font-semibold text-xl">
            <Image
              src={logo}
              alt="personal-image-Taker"
              className="invert w-10 h-10 mr-2"
            />
            Personal
            <span className="font-semibold text-xl text-accent ml-1">
              Note Taker.
            </span>
          </h1>
        </div>
      </Link>
      <div className="sm:hidden">
        <PopoverButton className="inline-flex items-center justify-center rounded-md  p-2 text-gray-400  hover:text-white focus:ring-2 focus:ring-inset">
          <Bars3Icon className="h-6 w-6" aria-hidden="true" />
        </PopoverButton>
      </div>
      <div className="hidden sm:flex gap-4">
        <Link
          href="/"
          className="text-lg rounded-full px-6 py-2 bg-transparent text-accent border border-accent hover:bg-accent hover:text-white cursor-pointer"
        >
          Home
        </Link>
        <Link
          href="/login"
          className="text-lg rounded-full px-6 py-2 bg-transparent text-accent border border-accent hover:bg-accent hover:text-white cursor-pointer"
        >
          Login
        </Link>
      </div>

      {/* Mobile Menu Panel */}
      <PopoverPanel className="fixed inset-y-0 right-0 z-10 w-64 bg-black text-white shadow-lg sm:hidden">
        <div className="p-4 flex flex-col h-full">
          <div className="flex justify-end mb-4">
            <PopoverButton className="inline-flex items-center justify-center rounded-md p-2 text-gray-400 hover:bg-gray-700 focus:ring-2 focus:ring-inset focus:ring-indigo-500">
              <XMarkIcon className="h-6 w-6" aria-hidden="true" />
            </PopoverButton>
          </div>
          <nav className="flex flex-col gap-6 mt-4">
            <Link
              href="/"
              className="text-lg rounded-full px-6 py-2 bg-transparent text-accent border border-accent hover:bg-accent hover:text-white cursor-pointer"
            >
              Home
            </Link>
            <Link
              href="/login"
              className="text-lg rounded-full px-6 py-2 bg-transparent text-accent border border-accent hover:bg-accent hover:text-white cursor-pointer"
            >
              Login
            </Link>
          </nav>
        </div>
      </PopoverPanel>
    </Popover>
  );
}
