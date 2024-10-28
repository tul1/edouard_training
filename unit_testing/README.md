# Assignment 16: Unit Testing with Mocking

## Objective

This assignment introduces unit testing in Go using the **testing** package, the **testify** library, and **mockgen** for generating mocks. You will implement tests for methods that depend on an interface, which you’ll mock to control the tests’ inputs and outputs.

## Setup Instructions

1. **Install Testify**
Ensure `testify` is installed by running

   ```bash
   go get github.com/stretchr/testify
   ```

2. **Install Mockgen**
Install the `mockgen` tool for generating mocks

   ```bash
   go install github.com/golang/mock/mockgen@v1.6.0
   ```

## Instructions

1. **Generate Mocks**:
   - Use `mockgen` to create a mock for the `DataService` interface.
   - Run the following command to generate the mock file in the `mocks` directory:

     ```bash
     mockgen -source=service.go -destination=mocks/mock_data_service.go -package=mocks
     ```

2. **Unit Tests**:
   - Create a file named `service_test.go`.
   - In this file, write unit tests for each method of the `Service` struct (`FetchData`, `UpdateData`, `FetchAndUpdateData`, `IsDataEqual`).
   - Use the generated mock (`mocks.MockDataService`) to simulate different behaviors and outcomes from the `DataService` dependency.

3. **Assertions**:
   - Use `require` and `assert` from `testify` for your assertions. `require` should be used for critical checks that prevent further test execution if failed (e.g., checking for no errors on essential calls), while `assert` is suitable for validating results that don’t impact the flow of the test.

## Example Test Cases

Here are some example test cases to guide you:

- **FetchData**:
  - Test fetching data successfully.
  - Test fetching data with an error response from `DataService.GetData`.

- **UpdateData**:
  - Test updating data successfully.
  - Test updating data with an error response from `DataService.SaveData`.

- **FetchAndUpdateData**:
  - Test fetching and updating data successfully.
  - Test a scenario where fetching data fails, ensuring `UpdateData` is not called.
  - Test a scenario where fetching is successful, but updating fails.

- **IsDataEqual**:
  - Test with matching data.
  - Test with non-matching data.
  - Test with an error response from `DataService.GetData`.

## Running the Tests

To run the tests, use the following command:

```bash
go test ./... -v
```

## Expected Output

The output will display pass/fail status for each test case, showing where assertions succeeded or failed.
