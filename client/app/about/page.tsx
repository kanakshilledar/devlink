const page = () => {
  const HLT = ({ children }: { children: React.ReactNode }) => {
    return <span className="text-white font-bold">{children}</span>;
  };

  const ContributorCard = ({ name, role }: { name: string; role: string }) => {
    return (
      <div className="flex flex-col gap-6 justify-center items-center border-2 rounded-lg p-6 min-w-64">
        <div className="w-44 h-48 bg-neutral-500" />
        <div className="flex flex-col items-center">
          <div className="text-2xl font-bold">{name}</div>
          <div className="text-xl text-neutral-400">{role}</div>
        </div>
      </div>
    );
  };
  return (
    <div className="flex justify-center">
      <div className="w-5/6 my-12">
        <div className="flex flex-col gap-6">
          <div className="text-5xl font-bold">About Us</div>
          <p className="text-2xl text-justify text-neutral-400">
            Welcome to <HLT>DevLink</HLT>, your go-to platform for discovering
            the best tech conferences, hackathons, and meetups happening around
            you. Whether you're a <HLT> tech enthusiast</HLT> looking to
            network, learn new skills, or showcase your talent, we aim to make
            it easy for you to find the events that matter to you. Our platform
            allows users to not only find events but also contribute by adding
            details of any tech events they discover. With a growing community,
            we aim to create a space where everyone can share opportunities,
            stay updated on upcoming events, and connect with like-minded people
            in the tech industry.
          </p>
        </div>
        <div className="flex flex-col gap-6 my-8">
          <div className="text-4xl font-bold">Maintainers:</div>
          <div className="flex items-center justify-center gap-6">
            <ContributorCard name="Ajay ram" role="Frontend Engineer" />
            <ContributorCard name="Kanak Shilledar" role="Backend Engineer" />
            <ContributorCard name="Gaurav" role="Backend Engineer" />
          </div>
        </div>
      </div>
    </div>
  );
};

export default page;
