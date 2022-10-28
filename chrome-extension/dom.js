setInterval(() => {
  document.querySelectorAll("article").forEach(function (el) {
    const prev = el.parentElement.innerHTML;
    if (!el.parentElement.innerHTML.includes("twitter-tags")) {
      el.innerHTML =
        `<div
      id="twitter-tags"
      style="
      color: white;
      min-height: 30px;
      width: 100%;
      font-size: 12px;
      display: flex;
      margin-top: 0.5rem;
      margin-left: 1rem;
      margin-right: 1rem;
      "
      >
      <div style="display:flex; margin-left:1rem;margin-right:1rem;">
      <div
        style="
        background-color: rgb(29 155 240);
        margin: 0.2rem;
        border-radius: 0.4rem;
        padding: 0.25rem;
      
        "
      >
        سایبری
      </div>
      <div style="display:flex; flex-direction: column; justify-content: center;font-size:9px;"><span>▲</span><span> ▼
      </span></div></div>
      <div
        style="
        background-color: rgb(29 155 240);
        margin: 0.2rem;
        border-radius: 0.4rem;
        padding: 0.25rem;
      
        "
      >
        قاسم چک
      </div>
      </div>
      
      ` + el.innerHTML;
    }
  });
}, 1000);
