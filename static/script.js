document.addEventListener("DOMContentLoaded", () => {
  document
    .getElementById("createUser")
    .addEventListener("submit", function (event) {
      event.preventDefault(); // Prevent the default form submission

      const form = event.target;
      const formData = new FormData(form);

      fetch("/users", {
        // Adjust the URL to match your routing
        method: "POST",
        body: formData,
      })
        .then((response) => response.json())
        .then((data) => {
          document.getElementById("responseMessage").innerText = data.message; // Display success message
        })
        .catch((error) => {
          document.getElementById("responseMessage").innerText =
            "Error creating user: " + error.message;
        });
    });

  const eventSource = new EventSource("/sse");
  eventSource.onmessage = function (e) {
    // let toast = document.getElementById("sse-data");

    // const test = setTimeout(() => {
    //   toast.classList.remove("hidden");
    // }, 3000);
    // test();
    console.log(e);
  };

  eventSource.onerror = function (error) {
    console.error("EventSource failed:", error);
    eventSource.close();
  };
  eventSource.onopen = function (e) {
    console.log("EventSource opened:", e);
  };
});
