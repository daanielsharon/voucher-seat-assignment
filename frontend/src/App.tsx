import {
  Alert,
  Button,
  Label,
  Modal,
  ModalBody,
  ModalHeader,
  Select,
  Spinner,
  TextInput,
} from "flowbite-react";
import { useEffect, useState } from "react";
import { check, generate } from "./api/apiClient";

export default function App() {
  const [crewName, setCrewName] = useState("");
  const [crewId, setCrewId] = useState("");
  const [flightNumber, setFlightNumber] = useState("");
  const [flightDate, setFlightDate] = useState("");
  const [aircraftType, setAircraftType] = useState("ATR");
  const [isSubmitAllowed, setIsSubmitAllowed] = useState(false);
  const [errorMessage, setErrorMessage] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const [seatResponse, setSeatResponse] = useState([]);

  function onSubmitSuccess() {
    setCrewName("");
    setCrewId("");
    setFlightNumber("");
    setFlightDate("");
    setAircraftType("ATR");
  }

  function setDefaultAlert() {
    setErrorMessage("");
    setSeatResponse([]);
  }

  function hasInputs(): boolean {
    if (!crewName || !crewId || !flightNumber || !flightDate || !aircraftType) {
      return false;
    }
    return true;
  }

  function onSubmit() {
    setIsLoading(true);
    setDefaultAlert();
    const checkBodyRequest = {
      flightNumber,
      date: flightDate,
    };

    const generateBodyRequest = {
      ...checkBodyRequest,
      name: crewName,
      id: crewId,
      aircraft: aircraftType,
    };

    check(checkBodyRequest)
      .then((res) => {
        if (res.data.exists) {
          setErrorMessage("Voucher already created");
          return;
        }
        generate(generateBodyRequest)
          .then((res) => {
            setSeatResponse(res.data.seats);
          })
          .catch((err) => {
            setErrorMessage(err.message);
          });
      })
      .catch((err) => {
        setErrorMessage(err.message);
      })
      .finally(() => {
        setIsLoading(false);
      });
  }

  useEffect(() => {
    if (hasInputs()) {
      setIsSubmitAllowed(true);
    } else {
      setIsSubmitAllowed(false);
    }
  }, [crewName, crewId, flightNumber, flightDate, aircraftType]);

  return (
    <>
      <Modal show={true} size="md" popup>
        <ModalHeader />
        <ModalBody>
          <div className="space-y-6">
            <h3 className="text-xl font-medium text-gray-900 dark:text-white">
              Voucher Seat Assignment
            </h3>
            {seatResponse.length > 0 && (
              <Alert
                color="success"
                onDismiss={() => {
                  onSubmitSuccess();
                  setSeatResponse([]);
                }}
              >
                <span className="font-medium">
                  Info seat: {seatResponse.join(", ")}
                </span>{" "}
                {`[Klik 'X' untuk reset form]`}
              </Alert>
            )}
            {!isSubmitAllowed && (
              <Alert color="warning" rounded>
                <span className="font-medium">{`[Info alert]`}</span> Please
                fill in all fields.
              </Alert>
            )}
            {errorMessage && (
              <Alert color="failure" rounded>
                <span className="font-medium">{`[Error alert]`}</span>{" "}
                {errorMessage}
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
              {/* <Datepicker value={flightDate} onChange={(event) => setFlightDate(event.target.value)} /> */}
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
              <Select
                defaultValue={undefined}
                id="aircraft"
                required
                value={aircraftType}
                onChange={(event) => setAircraftType(event.target.value)}
              >
                <option defaultChecked disabled value="" hidden>
                  Select Aircraft
                </option>
                <option>ATR</option>
                <option>Airbus 320</option>
                <option>Boeing 737 Max</option>
              </Select>
            </div>
            <div className="flex w-full justify-end">
              <Button
                disabled={!isSubmitAllowed || isLoading}
                onClick={onSubmit}
              >
                {isLoading ? (
                  <>
                    <Spinner
                      aria-label="Alternate spinner button example"
                      size="sm"
                    />
                    <span className="pl-3">Loading...</span>
                  </>
                ) : (
                  "Generate Vouchers"
                )}
              </Button>
            </div>
          </div>
        </ModalBody>
      </Modal>
    </>
  );
}
