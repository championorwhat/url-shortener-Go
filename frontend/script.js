const API_BASE_URL = "http://localhost:8080";

document.addEventListener("DOMContentLoaded", () => {
  document
    .getElementById("shortenBtn")
    .addEventListener("click", shortenUrl);
});

async function shortenUrl() {
  const input = document.getElementById("urlInput");
  const button = document.getElementById("shortenBtn");
  const resultDiv = document.getElementById("result");

  const longUrl = input.value.trim();

  if (!longUrl) {
    resultDiv.classList.remove("hidden");
    resultDiv.textContent = "Please enter a valid URL.";
    return;
  }

  button.disabled = true;
  button.textContent = "Shortening...";

  try {
    const response = await fetch(`${API_BASE_URL}/api/shorten`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ url: longUrl })
    });

    if (!response.ok) throw new Error("Request failed");

    const data = await response.json();
    const shortUrl = `${API_BASE_URL}/${data.short_code}`;

    resultDiv.classList.remove("hidden");
    resultDiv.innerHTML = `
      <div>
        <strong>Short URL:</strong><br/>
        <a href="${shortUrl}" target="_blank">${shortUrl}</a>
      </div>
      <button class="copy-btn" onclick="copyToClipboard('${shortUrl}')">
        Copy to Clipboard
      </button>
    `;
  } catch (err) {
    resultDiv.classList.remove("hidden");
    resultDiv.textContent = "Something went wrong. Try again.";
  } finally {
    button.disabled = false;
    button.textContent = "Shorten URL";
  }
}

function copyToClipboard(text) {
  navigator.clipboard.writeText(text).then(() => {
    alert("Short URL copied!");
  });
}
