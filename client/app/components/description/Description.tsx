import DescriptionImage from "../../../public/assests/descriptionImage.webp";
import Image from "next/image";
import AddnoteBtn from "../Homeaddnotebtn/Addnotebtn";
export default function Description() {
  return (
    <div className="flex flex-col md:flex-row items-center justify-center my-24 gap-8 md:gap-12 px-4">
      <div className="order-1 md:order-1 flex justify-center">
        <Image
          src={DescriptionImage}
          alt="DescriptionImage"
          className="w-full h-full md:w-full md:h-full object-cover"
        />
      </div>
      <div className="order-2 text-center md:text-left md:ml-12">
        <h1 className="text-xl md:text-4xl font-semibold">
          Why Choose Personal Note Taker?
        </h1>
        <p className="w-full md:w-[1000px] font-normal mt-4 text-md md:text-lg leading-relaxed">
          The Personal Note Taker application is your ultimate companion for
          organizing and storing notes related to all aspects of your life.
          Whether it's thoughts, plans, personal goals, or important topics, our
          app provides a secure and user-friendly platform to keep everything in
          one place. With features tailored for efficiency and convenience, you
          can categorize, access, and update your notes anytime, anywhere. Stay
          connected to your life's priorities and ensure no moment or idea is
          ever forgotten. Choose Personal Note Taker to simplify your life and
          keep your memories and plans organized seamlessly!
        </p>
        <div className=" order-3 md:order-3 flex mt-6 items-center justify-center ">
          <AddnoteBtn>Add Note </AddnoteBtn>
        </div>
      </div>
    </div>
  );
}
