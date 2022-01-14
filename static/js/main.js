const checkAvail = document.querySelector("#checkAvail");
checkAvail.addEventListener("click", () => {
    let form = `
    <form action="" class="login-address needs-validation" method="post" novalidate id="search-availability">
                    <div class="row">
                        <div class="col ">
                            <div class="row reservation-dates-modal">
                                <div class="col">
                                    <input disabled class="form-control" name="start" placeholder="Arrival date" required
                                           type="text" id="start" autocomplete="off">
                                </div>
                                <div class="col">
                                    <input disabled class="form-control" name="end" placeholder="Departure date" required
                                           type="text" id="end" autocomplete="off">
                                </div>
                            </div>
                        </div>
                    </div>
                </form>
                <hr>
    `
    attention.searchAvail({
        msg: form,
        title: "Search Availability",
        callback: function (result) {
            console.log("Jeremiah")
            let form = document.getElementById("search-availability")
            let formData = new FormData(form);
            formData.append("csrf_token", '{{.CSRFToken}}');

            fetch('/searchAvail-json', {
                method: "post",
                body: formData,
            }).then(response => response.json()).then(data => {
                console.log(data)
                console.log(data.ok)
                console.log(data.message)
            })
        }
    })
})
