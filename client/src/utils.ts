// Get categories
export async function getCategories() {
  return new Promise(async (resolve, reject) => {
    const response = await fetch("http://localhost:8000/categories", {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
    });
    const data = await response.json();
    if (data.status == "error" || !data.response) resolve(null)
    if (data.status == "success") resolve(data.response);
  });
}

export function checkIfValidNumberInput(event: any) {
  if (!event || !event.target) return;

  // Get the input box in which the event occured
  const target = event.target;
  const acceptedChars = "1234567890.";

  // Get the last inserted char
  const char = target.value.slice(-1);

  // Remove the last char from the input box, add it back if it passes all the tests
  target.value = target.value.slice(0, target.value.length - 1);

  // Check if the input char is one of the accepted chars, if else do nothing
  if (!(acceptedChars.includes(char))) return

  // Check if the entire value is 0, if so replace it with the input value
  if (target.value == "0") target.value = "";

  // Check if the input is a dot and there already exists a dot char, then do nothing
  if (char == "." && target.value.includes(".")) return;

  // Make it so if the input is empty and you add a dot, it will add a zero before it
  if (char == "." && !target.value.length) target.value = "0";

  // Check if the number is gonna start with a zero, if so do nothing
  if (target.value.slice(0, 2) == "0") return;

  // Add the value to the input box
  target.value += char;

}

export let CURRENCIES = ["RON", "USD", "EUR", "GBP"];
export let paymentMethods = ["VISA", "CASH", "MasterCard", "George"];
export let PORT = "8000"