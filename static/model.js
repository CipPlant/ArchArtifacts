document.addEventListener("DOMContentLoaded", function() {
    const form = document.getElementById("searchForm");
    const vek = document.getElementById("vek");
    const desyat = document.getElementById("desyat");
    const god = document.getElementById("god");
    const resultsContainer = document.getElementById("results-container");
    const interval = document.getElementById("interval");
    const interval2 = document.getElementById("interval2");
    const toggleSearchButton = document.getElementById("toggleSearchButton");
    const form2 = document.getElementById("intervalSearchForm");
    const intervalStart = document.getElementById("intervalStart");
    const intervalEnd = document.getElementById("intervalEnd");
    const resultItem = document.getElementById("toggleSearchButton");

    form2.addEventListener("submit", function(e) {
        e.preventDefault();
        const selectedInterval1 = intervalStart.value;
        const selectedInterval2 = intervalEnd.value;
        const url = new URL("/fuzzy", window.location.origin);
        url.searchParams.append("startValue", selectedInterval1);
        url.searchParams.append("endValue", selectedInterval2);
        fetch(url, {
            method: "POST"
        })
            .then(response => {
                if (response.status === 204) {
                    resultsContainer.style.backgroundColor = "rgba(200, 200, 200, 1)";
                    resultsContainer.style.display = "block";
                    resultsContainer.innerHTML = "Нет результатов";
                } else if (response.status === 400) {
                    response.text()
                        .then(errorMessage => {
                            resultsContainer.style.backgroundColor = "rgba(200, 200, 200, 1)";
                            resultsContainer.style.display = "block";
                            resultsContainer.innerHTML = errorMessage;
                        })
                        .catch(error => {
                            console.error("Ошибка при получении текста ошибки:", error);
                        });

                } else if (response.status === 500) {
                              resultsContainer.style.backgroundColor = "rgba(200, 200, 200, 1)";
                              resultsContainer.style.display = "block";
                              resultsContainer.innerHTML = "Ошибка сервера";
                } else if (response.status === 200) {
                     response.json()
                        .then(data => {
                            resultsContainer.innerHTML = "";
                            resultsContainer.style.backgroundColor = "rgba(200, 200, 200, 0)";
                            const ul = document.createElement("ul");
                            ul.style.display = "flex";
                            ul.style.flexWrap = "wrap";
                            ul.style.gap = "40px";
                            ul.style.justifyContent = "center";
                            ul.style.alignItems = "center";
                            data.forEach(result => {
                                const li = document.createElement("li");
                                li.className = "result-item";
                                const h2 = document.createElement("h2");
                                h2.textContent = result.buildingName;
                                li.appendChild(h2);
                                const image = document.createElement("img");
                                image.src = "data:image/jpeg;base64," + result.imagePath;
                                li.appendChild(image);
                                const textContainer = document.createElement("div");
                                textContainer.className = "text-container";
                                li.appendChild(textContainer);
                                const p = document.createElement("p");
                                p.textContent = result.buildingAge;
                                textContainer.appendChild(p);
                                const descr = addExpandButton(result.buildingDescription);
                                textContainer.appendChild(descr);
                                ul.appendChild(li);
                            });
                            resultsContainer.appendChild(ul);
                            resultsContainer.style.display = "block";

                        });
                } else {
                    console.error("Ошибка запроса:", response.status);
                }
            })
            .catch(error => {
                console.error("Ошибка запроса:", error);
            });
    });

    toggleSearchButton.addEventListener("click", () => {
        vek.value = "";
        const changeEvent = new Event("change", {
            bubbles: true,
            cancelable: false,
        });
        vek.dispatchEvent(changeEvent);
        resultsContainer.value = ""
        resultsContainer.style.display = "none";
        if (toggleSearchButton.classList.contains("right")) {
            toggleSearchButton.classList.remove("right");
            intervalSearch.style.display = "none";
            form.style.display = "block";
        } else {
            intervalStart.value = "";
            intervalEnd.value = ""
            toggleSearchButton.classList.add("right");
            intervalSearch.style.display = "block";
            form.style.display = "none";
        }
    });


    function updateDesyatOptions() {
        const selectedInterval = interval.value;
        const desyatOptions = desyat.getElementsByTagName("option");
        Array.from(desyatOptions).forEach((option, index) => {
            if (selectedInterval === "start") {
                option.hidden = ![1, 2, 3].includes(index);
            } else if (selectedInterval === "middle") {
                option.hidden = ![4, 5, 6, 7].includes(index);
            } else if (selectedInterval === "end") {
                option.hidden = ![8, 9, 10].includes(index);
            } else if (selectedInterval === "firstMiddle") {
                option.hidden = ![1, 2, 3, 4, 5].includes(index);
            } else if (selectedInterval === "secondMiddle") {
                option.hidden = ![6, 7, 8, 9, 10].includes(index);
            } else {
                option.hidden = false;
            }
        });
    }

    function updateGodOptions() {
        const selectedInterval = interval2.value;
        const godOptions = god.getElementsByTagName("option");
        Array.from(godOptions).forEach((option, index) => {
            if (selectedInterval === "start") {
                option.hidden = ![1, 2, 3].includes(index);
            } else if (selectedInterval === "middle") {
                option.hidden = ![4, 5, 6, 7].includes(index);
            } else if (selectedInterval === "end") {
                option.hidden = ![8, 9, 10].includes(index);
            } else if (selectedInterval === "firstMiddle") {
                option.hidden = ![1, 2, 3, 4, 5].includes(index);
            } else if (selectedInterval === "secondMiddle") {
                option.hidden = ![6, 7, 8, 9, 10].includes(index);
            } else {
                option.hidden = false;
            }
        });
    }
    interval2.addEventListener("change", function() {
        god.value = "";
        updateGodOptions()
    });
    interval.addEventListener("change", function() {
        desyat.value = "";
        god.value = "";
        interval2.value = "";
        updateDesyatOptions()
        updateGodOptions()
        god.setAttribute("disabled", true);
        interval2.setAttribute("disabled", true);
        if (vek.value !== "") {
            desyat.removeAttribute("disabled");
        } else {
            desyat.setAttribute("disabled", true);
        }
    });

    vek.addEventListener("change", function() {
        interval2.value = "";
        desyat.value = "";
        god.value = "";
        interval.value = "";
        god.setAttribute("disabled", true);
        interval2.setAttribute("disabled", true)
        interval.setAttribute("disabled", true)
        desyat.setAttribute("disabled", true);
        updateDesyatOptions()
        updateGodOptions()
        if (vek.value !== "") {
            desyat.removeAttribute("disabled");
            interval.removeAttribute("disabled");
        }
    });

    desyat.addEventListener("change", function() {
        god.value = "";
        interval2.value = ""
        updateGodOptions()
        if (desyat.value !== "") {
            interval2.removeAttribute("disabled");
            god.removeAttribute("disabled");
        } else {
            interval2.setAttribute("disabled", true);
            god.setAttribute("disabled", true);
        }
    });

    form.addEventListener("submit", function(e) {
        e.preventDefault();
        const selectedVek = vek.value;
        const selectedDesyat = desyat.value;
        const selectedGod = god.value;
        const selectedInterval1 = interval.value;
        const selectedInterval2 = interval2.value;
        const url = new URL("/search", window.location.origin);
        url.searchParams.append("vek", selectedVek);
        url.searchParams.append("desyat", selectedDesyat);
        url.searchParams.append("god", selectedGod);
        url.searchParams.append("interval1", selectedInterval1);
        url.searchParams.append("interval2", selectedInterval2);
        fetch(url, {
            method: "POST"
        })
            .then(response => {
                if (response.status === 204) {
                    resultsContainer.style.backgroundColor = "rgba(200, 200, 200, 1)";
                    resultsContainer.style.display = "block";
                    resultsContainer.innerHTML = "Нет результатов";
                } else if (response.status === 400) {
                    resultsContainer.style.backgroundColor = "rgba(200, 200, 200, 1)";
                    resultsContainer.style.display = "block";
                    resultsContainer.innerHTML = "Ошибка в запросе";
                } else if (response.status === 500) {
                      resultsContainer.style.backgroundColor = "rgba(200, 200, 200, 1)";
                      resultsContainer.style.display = "block";
                      resultsContainer.innerHTML = "Ошибка сервера";
                 } else if (response.status === 200) {
                    response.json()
                        .then(data => {
                            resultsContainer.innerHTML = "";
                            resultsContainer.style.backgroundColor = "rgba(200, 200, 200, 0)";
                            const ul = document.createElement("ul");
                            ul.style.display = "flex";
                            ul.style.flexWrap = "wrap";
                            ul.style.gap = "40px";
                            ul.style.justifyContent = "center";
                            ul.style.alignItems = "center";
                            data.forEach(result => {
                                const li = document.createElement("li");
                                li.className = "result-item";
                                li.style.width = "100%";
                                const h2 = document.createElement("h2");
                                h2.textContent = result.buildingName;
                                li.appendChild(h2);
                                const image = document.createElement("img");
                                image.src = "data:image/jpeg;base64," + result.imagePath;
                                li.appendChild(image);
                                const textContainer = document.createElement("div");
                                textContainer.className = "text-container";
                                li.appendChild(textContainer);
                                const p = document.createElement("p");
                                p.textContent = result.buildingAge;
                                textContainer.appendChild(p);
                                const descr = addExpandButton(result.buildingDescription);
                                textContainer.appendChild(descr);
                                ul.appendChild(li);
                            });
                            resultsContainer.appendChild(ul);
                            resultsContainer.style.display = "block";
                        });
                } else {
                    console.error("Ошибка запроса:", response.status);
                }
            })
            .catch(error => {
                console.error("Ошибка запроса:", error);
            });
    });
});

function addExpandButton(description) {
    const maxCharacters = 100;
    const descr = document.createElement('descr');
    descr.textContent = description;
    if (description.length > maxCharacters) {
        const truncatedText = description.slice(0, maxCharacters) + '...'
        const expandButton = document.createElement('button');
        expandButton.textContent = 'Развернуть';
        expandButton.className = 'expand-button';
        expandButton.addEventListener('click', function () {
            expandButton.style.display = 'none';
            descr.textContent = description;
        });
        descr.textContent = truncatedText;
        descr.appendChild(expandButton);
    }
    return descr;
}
