import {
  Alert,
  Button,
  Label,
  Modal,
  ModalBody,
  ModalHeader,
  Select,
  TextInput,
} from "flowbite-react";
import { useEffect, useState } from "react";
import { check, generate } from "./api/apiClient";

export default function App() {
  const [crewName, setCrewName] = useState("");
  const [crewId, setCrewId] = useState("");
  const [flightNumber, setFlightNumber] = useState("");
  const [flightDate, setFlightDate] = useState("");
  const [aircraftType, setAircraftType] = useState("")
  const [isSubmitAllowed, setIsSubmitAllowed] = useState(false);
  const [errorMessage, setErrorMessage] = useState("")

  function onSubmitSuccess() {
    setCrewName("");
    setCrewId("");
    setFlightNumber("");
    setFlightDate("");
    setAircraftType("");
  }

  function hasInputs(): boolean {
    if (!crewName || !crewId || !flightNumber || !flightDate || !aircraftType) {
      return false;
    }
    return true;
  }

  function onSubmit() {
    const checkBodyRequest = {
      flightNumber,
      date: flightDate
    }

    check(checkBodyRequest)
    .then((res) => {
      console.log(res)
    })
    .catch((err) => {
      setErrorMessage(err.message)
    })

    const generateBodyRequest = {
      ...checkBodyRequest,
      name: crewName,
      id: crewId,
      aircraft: aircraftType
    }

    generate(generateBodyRequest)
    .then((res) => {
      console.log(res)
    })
    .catch((err) => {
      setErrorMessage(err.message)
    })

    onSubmitSuccess()
  }

  useEffect(() => {
    if (hasInputs()) {
      setIsSubmitAllowed(true);
    } else {
      setIsSubmitAllowed(false);
    }
  }, [crewName, crewId, flightNumber, flightDate, aircraftType]);

  return (
    <Modal show={true} size="md" popup>
      <ModalHeader />
      <ModalBody>
        <div className="space-y-6">
          <h3 className="text-xl font-medium text-gray-900 dark:text-white">
            Voucher Seat Assignment
          </h3>
          {!isSubmitAllowed && (
            <Alert color="warning" rounded>
              <span className="font-medium">{`[Info alert]`}</span> Please fill
              in all fields.
            </Alert>
          )}
          {isSubmitAllowed && errorMessage && (
            <Alert color="failure" rounded>
              <span className="font-medium">{`[Error alert]`}</span> {errorMessage}
            </Alert>
          )}
          <div>
            <div className="mb-2 block">
              <Label htmlFor="crew_name">Crew Name</Label>
            </div>
            <TextInput
              id="crew_name"
              type="text"
              placeholder="Your name"
              value={crewName}
              onChange={(event) => setCrewName(event.target.value)}
              required
            />
          </div>
          <div>
            <div className="mb-2 block">
              <Label htmlFor="crew_id">Crew ID</Label>
            </div>
            <TextInput
              id="crew_id"
              type="text"
              placeholder="Your ID"
              value={crewId}
              onChange={(event) => setCrewId(event.target.value)}
              required
            />
          </div>
          <div>
            <div className="mb-2 block">
              <Label htmlFor="flight_number">Flight Number</Label>
            </div>
            <TextInput
              id="flight_number"
              type="text"
              placeholder="Flight Number"
              value={flightNumber}
              onChange={(event) => setFlightNumber(event.target.value)}
              required
            />
          </div>
          <div>
            <div className="mb-2 block">
              <Label htmlFor="flight_date">Flight Date</Label>
            </div>
            <TextInput
              id="flight_date"
              type="date"
              placeholder="Flight Date"
              value={flightDate}
              onChange={(event) => setFlightDate(event.target.value)}
              required
            />
          </div>
          <div className="max-w-md">
            <div className="mb-2 block">
              <Label htmlFor="aircraft">Select Aircraft</Label>
            </div>
            <Select id="aircraft" required value={aircraftType} onChange={(event) => setAircraftType(event.target.value)}>
              <option defaultChecked disabled value="" hidden>
                Select Aircraft
              </option>
              <option>ATR</option>
              <option>Airbus 320</option>
              <option>Boeing 737 Max</option>
            </Select>
          </div>
          <div className="flex w-full justify-end">
            <Button disabled={!isSubmitAllowed} onClick={onSubmit}>Generate Vouchers</Button>
          </div>
        </div>
      </ModalBody>
    </Modal>
  );
}
