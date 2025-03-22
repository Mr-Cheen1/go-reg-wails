import { useEffect, useState } from "react";
import { UpdateProduct } from "../../wailsjs/go/main/App";
import { models } from "../../wailsjs/go/models";
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

type Product = models.Product;

interface EditProductDialogProps {
  isOpen: boolean;
  onClose: () => void;
  product: Product;
  onSuccess: () => void;
}

export function EditProductDialog({
  isOpen,
  onClose,
  product,
  onSuccess,
}: EditProductDialogProps) {
  const [name, setName] = useState("");
  const [timeCalculation, setTimeCalculation] = useState("");
  const [isSubmitting, setIsSubmitting] = useState(false);
  const { toast } = useToast();

  useEffect(() => {
    if (product) {
      setName(product.name);
      setTimeCalculation(product.timeCalculation);
    }
  }, [product]);

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
      await UpdateProduct(product.id, name, timeCalculation);
      toast({
        title: "Успешно",
        description: "Запись обновлена",
      });
      onSuccess();
      onClose();
    } catch (error) {
      toast({
        title: "Ошибка",
        description: "Не удалось обновить запись",
        variant: "destructive",
      });
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <Dialog open={isOpen} onOpenChange={onClose}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Редактировать запись</DialogTitle>
        </DialogHeader>

        <div className="grid gap-4 py-4">
          <div className="grid gap-2">
            <label htmlFor="edit-name" className="text-sm font-medium">
              Наименование
            </label>
            <Input
              id="edit-name"
              value={name}
              onChange={(e) => setName(e.target.value)}
            />
          </div>
          <div className="grid gap-2">
            <label htmlFor="edit-time" className="text-sm font-medium">
              Время обработки
            </label>
            <Input
              id="edit-time"
              value={timeCalculation}
              onChange={(e) => setTimeCalculation(e.target.value)}
              placeholder="Например: 8+2+5"
            />
          </div>
        </div>

        <DialogFooter>
          <Button variant="outline" onClick={onClose}>
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