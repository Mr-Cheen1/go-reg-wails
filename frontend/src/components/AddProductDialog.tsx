import { useState } from "react";
import { AddProduct } from "../../wailsjs/go/main/App";
import { Button } from "./ui/button";
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "./ui/dialog";
import { Input } from "./ui/input";
import { useToast } from "../hooks/use-toast";

interface AddProductDialogProps {
  isOpen: boolean;
  onClose: () => void;
  onSuccess: () => void;
}

export function AddProductDialog({
  isOpen,
  onClose,
  onSuccess,
}: AddProductDialogProps) {
  const [name, setName] = useState("");
  const [timeCalculation, setTimeCalculation] = useState("");
  const [isSubmitting, setIsSubmitting] = useState(false);
  const { toast } = useToast();

  const handleSubmit = async () => {
    if (!name.trim()) {
      toast({
        title: "Ошибка",
        description: "Наименование не может быть пустым",
        variant: "destructive",
      });
      return;
    }

    setIsSubmitting(true);
    try {
      await AddProduct(name, timeCalculation);
      toast({
        title: "Успешно",
        description: "Запись добавлена",
      });
      onSuccess();
      handleClose();
    } catch (error) {
      toast({
        title: "Ошибка",
        description: "Не удалось добавить запись",
        variant: "destructive",
      });
    } finally {
      setIsSubmitting(false);
    }
  };

  const handleClose = () => {
    setName("");
    setTimeCalculation("");
    onClose();
  };

  return (
    <Dialog open={isOpen} onOpenChange={handleClose}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Добавить запись</DialogTitle>
        </DialogHeader>

        <div className="grid gap-4 py-4">
          <div className="grid gap-2">
            <label htmlFor="name" className="text-sm font-medium">
              Наименование
            </label>
            <Input
              id="name"
              value={name}
              onChange={(e) => setName(e.target.value)}
              placeholder="Введите наименование"
            />
          </div>
          <div className="grid gap-2">
            <label htmlFor="timeCalculation" className="text-sm font-medium">
              Время обработки
            </label>
            <Input
              id="timeCalculation"
              value={timeCalculation}
              onChange={(e) => setTimeCalculation(e.target.value)}
              placeholder="Например: 8+2+5"
            />
          </div>
        </div>

        <DialogFooter>
          <Button variant="outline" onClick={handleClose}>
            Отмена
          </Button>
          <Button onClick={handleSubmit} disabled={isSubmitting}>
            {isSubmitting ? "Сохранение..." : "Сохранить"}
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
} 