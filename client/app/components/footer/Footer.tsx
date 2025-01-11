import Link from "next/link";
import Image from "next/image";

import facebookIcon from "../../../public/assests/Facebook.png";
import instagramIcon from "../../../public/assests/insta.png";
import twitterIcon from "../../../public/assests/Twitter.png";

export default function Footer() {
  return (
    <div className="bg-black py-8">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
          <div className="space-y-4 text-center md:text-left">
            <h2 className="text-md md:text-xl font-semibold text-white">
              Quick Links
            </h2>
            <div className=" space-x-3 md:space-x-5">
              <Link
                href="/home"
                className="text-base text-gray-300 hover:text-accent transition duration-300"
              >
                Home
              </Link>
              <Link
                href="/login"
                className="text-base text-gray-300 hover:text-accent transition duration-300"
              >
                Login
              </Link>
            </div>
          </div>
          <div className="flex flex-col items-center text-center">
            <p className="text-lg text-gray-300">
              We strive to provide the best service for you.
            </p>
          </div>
          <div className="space-y-4 text-center">
            <h2 className="text-xl font-semibold text-accent">
              Connect with Us
            </h2>
            <div className="flex justify-center gap-6">
              <Image
                src={facebookIcon}
                alt="Facebook"
                className="w-6 h-6 transition duration-300 filter brightness-0 invert hover:brightness-100 cursor-pointer"
              />
              <Image
                src={instagramIcon}
                alt="Instagram"
                className="w-6 h-6 transition duration-300 filter brightness-0 invert hover:brightness-100 cursor-pointer"
              />
              <Image
                src={twitterIcon}
                alt="Twitter"
                className="w-6 h-6 transition duration-300 cursor-pointer"
              />
            </div>
          </div>
        </div>
        <div className="mt-8 border-t border-gray-700 pt-6 text-center w-full">
          <p className="text-sm text-gray-400 leading-tight">
            &copy; {new Date().getFullYear()} Personal Note Taker. All Rights
            Reserved.
          </p>
        </div>
      </div>
    </div>
  );
}
