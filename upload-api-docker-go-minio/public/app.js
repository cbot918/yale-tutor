// Function to handle the file upload via Fetch API
async function uploadFile(file) {
  const formData = new FormData();
  formData.append('file', file);

  try {
      const response = await fetch('http://localhost:3456/upload', {
          method: 'POST',
          body: formData,
      });

      if (!response.ok) {
          throw new Error(`Error: ${response.statusText}`);
      }

      const result = await response.text();
      alert(`Success: ${result}`);
  } catch (error) {
      alert(`Failed to upload: ${error.message}`);
  }
}

// Add event listener to the form
document.getElementById('uploadForm').addEventListener('submit', function(event) {
  event.preventDefault(); // Prevent the default form submission

  const fileInput = document.getElementById('fileInput');
  const file = fileInput.files[0]; // Get the file from the input

  if (file) {
      uploadFile(file); // Upload the file
  } else {
      alert('Please select a file to upload.');
  }
});

document.getElementById('ping').addEventListener('click', function(event) {
  event.preventDefault(); // Prevent the default form submission

  fetch("http://localhost:3456/ping")
    // .then(res=>res.json())
    .then(data=>console.log(data))
    .catch(err=>console.log(err))
});

document.getElementById('testpost').addEventListener('click', function(event) {
  event.preventDefault(); // Prevent the default form submission

  fetch("http://localhost:3456/post",{
    method: 'POST',
    headers: {
      'Content-Type': 'application/json', // Specify content type if you're sending JSON data
      // For CORS issues, you might not need to set 'Content-Type' if not sending data
      // 'Content-Type': 'application/x-www-form-urlencoded',
    },
  })
    // .then(res=>res.json())
    .then(data=>console.log(data))
    .catch(err=>console.log(err))
});