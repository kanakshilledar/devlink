import Link from "next/link";

const page = () => {
  return (
    <div className="flex justify-center w-full">
      <div className="w-4/5 my-8">
        <Link href="/">
          <div className="font-bold text-3xl hover:underline cursor-pointer text-neutral-400">
            &lt; Home
          </div>
        </Link>
        <div className="text-5xl font-bold mt-4">Terms of Use:</div>
        <div className="my-6">
          <div className="text-xl font-bold">1. Acceptance of Terms</div>
          <p>
            By accessing and using DevLink, you agree to comply with and be
            bound by these Terms of Use. If you do not agree to these terms, you
            should not use this Website.
          </p>
        </div>
        <div className="my-6">
          <div className="text-xl font-bold">2. User Responsibilities</div>
          <p>
            As a user of the Website, you agree to provide accurate and truthful
            information when contributing event details. You are solely
            responsible for the content you submit, including ensuring that it
            does not violate any third-party rights or applicable laws.
          </p>
        </div>
        <div className="my-6">
          <div className="text-xl font-bold">3. Event Submission</div>
          <p>
            Users can submit information about tech events (conferences,
            hackathons, meetups, etc.). By submitting an event, you agree that:
          </p>
          <ul>
            <li>
              The event details provided are accurate to the best of your
              knowledge.
            </li>
            <li>You have the right to share the event information.</li>
            <li>
              DevLink is not responsible for verifying the accuracy or validity
              of submitted events.
            </li>
          </ul>
        </div>
        <div className="my-6">
          <div className="text-xl font-bold">
            4. Content Ownership and License
          </div>
          <p>
            All content you submit to the Website remains your intellectual
            property. However, by submitting content, you grant DevLink a
            worldwide, non-exclusive, royalty-free license to use, distribute,
            and display the content for purposes related to the operation and
            promotion of the Website.
          </p>
        </div>
        <div className="my-6">
          <div className="text-xl font-bold">
            5. Intellectual Property Rights
          </div>
          <p>
            The content and design of the Website, including logos, text,
            graphics, and software, are the property of DevLink and are
            protected by intellectual property laws. Users are prohibited from
            copying, distributing, or otherwise using any part of the Website
            without prior written consent.
          </p>
        </div>
        <div className="my-6">
          <div className="text-xl font-bold">6. Disclaimer of Warranties</div>
          <p>
            The Website is provided "as is" without any warranties, either
            express or implied. We do not guarantee the accuracy, reliability,
            or completeness of any information on the Website, including event
            details. We are not responsible for any damages resulting from the
            use of or inability to use the Website.
          </p>
        </div>
        <div className="my-6">
          <div className="text-xl font-bold">7. Limitation of Liability</div>
          <p>
            In no event shall DevLink, its owners, or its contributors be liable
            for any direct, indirect, incidental, or consequential damages
            arising out of your use of the Website.
          </p>
        </div>
        <div className="my-6">
          <div className="text-xl font-bold">8. User Conduct</div>
          <p>
            You agree not to engage in any of the following prohibited
            activities:
          </p>
          <ul>
            <li>Violating any laws or regulations.</li>
            <li>Submitting false or misleading event information.</li>
            <li>Posting harmful, defamatory, or abusive content.</li>
            <li>
              Attempting to interfere with the security or operation of the
              Website.
            </li>
          </ul>
        </div>
        <div className="my-6">
          <div className="text-xl font-bold">9. Modification of Terms</div>
          <p>
            DevLink reserves the right to modify these Terms of Use at any time.
            Users will be notified of any significant changes, and continued use
            of the Website following such changes will constitute your
            acceptance of the new terms.
          </p>
        </div>
        <div className="my-6">
          <div className="text-xl font-bold">10. Termination</div>
          <p>
            We reserve the right to suspend or terminate your access to the
            Website for any reason, including violation of these Terms of Use.
          </p>
        </div>
        <div className="my-6">
          <div className="text-xl font-bold">11. Governing Law</div>
          <p>
            These Terms of Use shall be governed by and construed in accordance
            with the laws of India. Any disputes arising under these terms shall
            be subject to the exclusive jurisdiction of the courts in Chennai
          </p>
        </div>
        <div className="my-6">
          <div className="text-xl font-bold">12. Contact Information</div>
          <p>
            If you have any questions or concerns about these Terms of Use,
            please contact us at [Email Address].
          </p>
        </div>
      </div>
    </div>
  );
};

export default page;
