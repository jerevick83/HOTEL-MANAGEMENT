// Fetch all the forms we want to apply custom Bootstrap validation styles to
const forms = document.querySelectorAll('.needs-validation');
const sweetAlertModule = document.querySelector("#sub");
let attention = prompt();
const checkAvail = document.querySelector("#checkAvail");
(function () {
    'use strict'
    // Loop over the forms and prevent submission
    Array.prototype.slice.call(forms)
        .forEach(function (form) {
            form.addEventListener('submit', function (event) {
                if (!form.checkValidity()) {
                    event.preventDefault()
                    event.stopImmediatePropagation()

                }

                form.classList.add('was-validated')

            }, false)

        })

})()

checkAvail.addEventListener("click", () => {
    let form = `
    <form action="reservation.html" class="login-address needs-validation" method="GET" novalidate>
                    <div class="row">
                        <div class="col ">
                            <div class="row reservation-dates-modal "  >
                                <div class="col">
                                    <input disabled class="form-control" name="start" placeholder="Arrival date" required
                                           type="text" id="start" >
                                </div>
                                <div class="col">
                                    <input disabled class="form-control" name="end" placeholder="Departure date" required
                                           type="text" id="end">
                                </div>
                            </div>
                        </div>
                    </div>
<!--                    <button class="btn btn-primary mt-3" type="submit">Search Availability</button>-->
                </form>
                <hr>
    `
    attention.searchAvail({msg: form, title: "Search Availability"})
})

function prompt() {
    let success = (c) => {

        const {
            msg = "",
            icon = "success",
            position = "top-end"
        } = c;
        const Toast = Swal.mixin({
            toast: true,
            title: msg,
            icon: icon,
            position: position,
            showCancelButton: false,
            timer: 3000,
            timerProgressBar: true,
            didOpen: (toast) => {
                toast.addEventListener('mouseenter', Swal.stopTimer)
                toast.addEventListener('mouseleave', Swal.resumeTimer)
            }
        })

        Toast.fire({})
    };
    let warning = () => {
        console.log("Hello From warning")
    };
    let error = () => {
        console.log("This is an error")
    };

    // searchAvail check for room availability
    let searchAvail = async (c) => {
        const {
            msg = "",
            title = ""
        } = c
        const {value: formValues} = await Swal.fire({
            title: title,
            html: msg,
            backdrop: false,
            showCancelButton: true,
            customClass: {
                cancelButton: "cancelBtn",
                confirmButton: "confirmButton"
            },
            focusConfirm: false,
            width: '35rem',
            scrollbarPadding: false,
            inputAutoTrim: true,
            buttonsStyling: false,
            willOpen: () => {
                const calendarPicker = document.querySelector(".reservation-dates-modal");
                new DateRangePicker(calendarPicker, {
                    format: "yyyy-mm-dd",
                    orientation: "top",
                    todayBtn: true,
                    todayBtnMode: 1,
                    autohide: true,
                    allowOneSidedRange: true,
                    showOnFocus: true
                });
            },
            preConfirm: () => {
                return [
                    document.getElementById('start').value,
                    document.getElementById('end').value
                ]
            },
            didOpen: () => {
                document.getElementById('start').removeAttribute("disabled");
                document.getElementById('end').removeAttribute("disabled");
            }
        })

        if (formValues) {
            Swal.fire(JSON.stringify(formValues))
        }
    }
    return ({
            success,
            warning,
            error,
            searchAvail
        }
    )
}

//attention is receiving the returned methods from the prompt function

// this is the success method from the prompt function
// sweetAlertModule.addEventListener("click", ()=>attention.success({msg: "Thanks for signing in"}));

// this is the search availability method from the prompt function

// Calendar Picker
